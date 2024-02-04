package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	"github.com/skip2/go-qrcode"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"strconv"
	"strings"
	"time"
)

type CheckExistingQuotationStruct struct {
	*models.Quotation
}

func (module *module) RetrieveAllQuotation(id uuid.UUID) (m []models.Quotation, err error) {
	if m, err = module.db.quotationModel.GetAllQuotation(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllQuotationPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.quotationModel.GetAllQuotationPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveQuotation(id uuid.UUID) (m models.Quotation, err error) {
	if m, err = module.db.quotationModel.GetOneQuotationByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertQuotation(p *dto.InsertQuotation) (err error) {

	tx := database.GetDatabaseConnection()
	QuotationModel := models.Quotation{}

	_, quotationLastDataErr := database.GetLastDocumentNumber(tx, &QuotationModel)

	var formattedNumber string
	var finalNumber string
	if quotationLastDataErr != nil {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", "0001", "SPH", monthInRoman, year)
		finalNumber = "0001"
	} else {
		parsedNumber, _ := strconv.Atoi(QuotationModel.Sequence)
		strLastNumber := strconv.Itoa(parsedNumber)
		getFrontDigitNumber := strings.Replace(QuotationModel.Sequence, strLastNumber, "", -1)
		convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
		addActiveNumberByOne := convertActiveNumberToStr + 1
		convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
		poNumber := getFrontDigitNumber + convertAddedActiveNumberToStr

		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", poNumber, "SPH", monthInRoman, year)
		finalNumber = poNumber
	}

	if p.OperatingActivityID.String() != "" {
		Quotation := models.Quotation{
			Number:              formattedNumber,
			FrancoArea:          p.FrancoArea,
			PaymentTerm:         p.PaymentTerm,
			SendAfter:           p.SendAfter,
			Date:                p.Date,
			Sequence:            finalNumber,
			OperatingActivityID: p.OperatingActivityID,
			QuotationCreatedBy:  p.CreatedBy,
		}

		if QuotationErr := tx.Create(&Quotation); err != nil {
			tx.Rollback()
			return QuotationErr.Error
		}
	} else {
		Quotation := models.Quotation{
			Number:             formattedNumber,
			FrancoArea:         p.FrancoArea,
			PaymentTerm:        p.PaymentTerm,
			SendAfter:          p.SendAfter,
			Date:               p.Date,
			Sequence:           finalNumber,
			QuotationCreatedBy: p.CreatedBy,
		}

		if QuotationErr := tx.Create(&Quotation); err != nil {
			tx.Rollback()
			return QuotationErr.Error
		}
	}
	tx.Commit()
	return
}

func (module *module) UpdateQuotation(id uuid.UUID, p *dto.UpdateQuotation) (err error) {

	if p.OperatingActivityID.String() != "" {
		if err = module.db.quotationModel.UpdateQuotation(id, models.Quotation{
			FrancoArea:          p.FrancoArea,
			PaymentTerm:         p.PaymentTerm,
			SendAfter:           p.SendAfter,
			Date:                p.Date,
			OperatingActivityID: p.OperatingActivityID,
			QuotationUpdatedBy:  p.UpdatedBy,
		}); err != nil {
			return errors.New(err.Error())
		}
	} else {
		if err = module.db.quotationModel.UpdateQuotation(id, models.Quotation{
			FrancoArea:         p.FrancoArea,
			PaymentTerm:        p.PaymentTerm,
			SendAfter:          p.SendAfter,
			Date:               p.Date,
			QuotationUpdatedBy: p.UpdatedBy,
		}); err != nil {
			return errors.New(err.Error())
		}
	}
	return
}

func (module *module) GenerateSPHDocument(id uuid.UUID) (output GenerateDocumentOutput, err error) {

	formatCurrency := accounting.Accounting{Symbol: "", Precision: 2}

	var quotation models.Quotation

	if quotation, err = module.RetrieveQuotation(id); err != nil {
		return output, errors.New("quotation not found")
	}

	var subTotalProduct float64
	var grandTotalProduct string

	var remappedProduct []RemappedProduct

	if len(quotation.OperatingActivity.OperatingActivityProduct) > 0 {
		for _, data := range quotation.OperatingActivity.OperatingActivityProduct {
			subTotalProduct += data.SubTotal

			subTotalProductFormatted := formatCurrency.FormatMoney(data.SubTotal)
			unitPriceProductFormatted := formatCurrency.FormatMoney(data.Product.UnitPrice * 0.11)
			grandTotalProduct = formatCurrency.FormatMoney(data.SubTotal + (data.SubTotal * 0.11))

			productStruct := models.CustomProduct{
				Name:      data.Product.Name,
				UnitPrice: unitPriceProductFormatted,
				Packaging: data.Product.Packaging,
				Stock:     data.Product.Stock,
				Note:      data.Product.Note,
			}
			productData := RemappedProduct{
				Quantity:   data.Quantity,
				VATRate:    data.VATRate,
				SubTotal:   subTotalProductFormatted,
				GrandTotal: grandTotalProduct,
				Product:    productStruct,
			}
			remappedProduct = append(remappedProduct, productData)
		}
	} else {
		return output, errors.New("cannot generate if there is no product in this operating activity product")
	}

	png, err := qrcode.Encode(quotation.Number, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	dataURI := base64.StdEncoding.EncodeToString(png)

	templateData := struct {
		ClientName     string
		DocumentNumber string
		Products       []RemappedProduct
		QrCode         string
	}{
		DocumentNumber: quotation.Number,
		ClientName:     quotation.OperatingActivity.Client.Name,
		Products:       remappedProduct,
		QrCode:         dataURI,
	}

	var outputPath string
	var saveFileErr error
	if outputPath, saveFileErr = utils.SaveFileToDockerVolume(nil, "", "quotation", nil, templateData); saveFileErr != nil {
		return output, errors.New(saveFileErr.Error())
	}

	data := GenerateDocumentOutput{
		OutputPath: outputPath,
		FileName:   quotation.Number,
	}

	return data, nil
}

func (module *module) CheckExistingQuotation(id string, param CheckExistingQuotationStruct) (err error) {

	if param.Number != "" {
		if _, dbErr := module.db.quotationModel.GetOneQuotationByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.quotationModel.GetOneQuotationByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
