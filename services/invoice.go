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

type CheckExistingInvoiceStruct struct {
	*models.Invoice
}

func (module *module) RetrieveInvoice(id uuid.UUID) (m models.Invoice, err error) {
	if m, err = module.db.invoiceModel.GetOneInvoiceByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllInvoice(id uuid.UUID) (m []models.Invoice, err error) {
	if m, err = module.db.invoiceModel.GetAllInvoice(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) GenerateInvoiceDocument(id uuid.UUID) (output GenerateDocumentOutput, err error) {

	formatCurrency := accounting.Accounting{Symbol: "", Precision: 2}

	var invoice models.Invoice

	if invoice, err = module.RetrieveInvoice(id); err != nil {
		return output, errors.New("invoice not found")
	}

	var subTotalProduct float64
	var grandTotalProduct string
	var ppnProduct string

	var remappedProduct []RemappedProduct

	if len(invoice.OperatingActivity.OperatingActivityProduct) > 0 {
		for _, data := range invoice.OperatingActivity.OperatingActivityProduct {
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

	subTotalProductFormatted := formatCurrency.FormatMoney(subTotalProduct)
	ppnProduct = formatCurrency.FormatMoney(subTotalProduct * 0.11)
	grandTotalProduct = formatCurrency.FormatMoney(subTotalProduct + (subTotalProduct * 0.11))
	date := utils.ConvertEnToIDDateTime(invoice.Date)

	png, err := qrcode.Encode(invoice.Number, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	var purchaseOrder models.PurchaseOrder
	if purchaseOrder, err = module.db.purchaseOrderModel.GetOnePurchaseOrderByOperatingActivityId(invoice.OperatingActivityID); err != nil {
		return output, fmt.Errorf("purchase order %s", err.Error())
	}

	var deliveryOrder models.DeliveryOrder
	if deliveryOrder, err = module.db.deliveryOrderModel.GetOneDeliveryOrderByOperatingActivityID(invoice.OperatingActivityID); err != nil {
		return output, fmt.Errorf("delivery order %s", err.Error())
	}

	shippingDate := utils.ConvertEnToIDDateTime(deliveryOrder.Date)

	var payment models.Payment
	if payment, err = module.db.paymentModel.GetOnePaymentByOperatingActivityID(invoice.OperatingActivityID); err != nil {
		return output, fmt.Errorf("payment %s", err.Error())
	}

	dataURI := base64.StdEncoding.EncodeToString(png)

	templateData := struct {
		ClientName     string
		ClientAddress  string
		DocumentNumber string
		PONumber       string
		DONumber       string
		PaymentTerm    string
		ShippingDate   string
		DateGenerated  string
		Products       []RemappedProduct
		SubTotal       string
		Ppn            string
		GrandTotal     string
		QrCode         string
	}{
		ClientName:     invoice.OperatingActivity.Client.Name,
		ClientAddress:  invoice.OperatingActivity.Client.Address,
		DocumentNumber: invoice.Number,
		PONumber:       purchaseOrder.Number,
		DONumber:       deliveryOrder.Number,
		PaymentTerm:    payment.Term,
		ShippingDate:   shippingDate,
		DateGenerated:  date,
		Products:       remappedProduct,
		SubTotal:       subTotalProductFormatted,
		Ppn:            ppnProduct,
		GrandTotal:     grandTotalProduct,
		QrCode:         dataURI,
	}

	var outputPath string
	var saveFileErr error
	if outputPath, saveFileErr = utils.SaveFileToDockerVolume(nil, "", "invoice", nil, templateData); saveFileErr != nil {
		return output, errors.New(saveFileErr.Error())
	}

	data := GenerateDocumentOutput{
		OutputPath: outputPath,
		FileName:   invoice.Number,
	}

	return data, nil
}

func (module *module) RetrieveAllInvoicePaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.invoiceModel.GetAllInvoicePaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertInvoice(p *dto.InsertInvoice) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	InvoiceModel := models.Invoice{}

	invoiceType := strings.ToLower(p.Type)

	_, invoiceLastDataErr := database.GetLastDocumentInvoiceNumber(invoiceType, tx, &InvoiceModel)

	var formattedNumber string
	var finalNumber string
	if invoiceLastDataErr != nil {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		finalNumber = "0001"
		if p.Type == "proforma" {
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "PRF", monthInRoman, year)
		}

		if p.Type == "invoice" {
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "INV", monthInRoman, year)
		}
	} else {
		parsedNumber, _ := strconv.Atoi(InvoiceModel.Sequence)
		strLastNumber := strconv.Itoa(parsedNumber)
		getFrontDigitNumber := strings.Replace(InvoiceModel.Sequence, strLastNumber, "", -1)
		convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
		addActiveNumberByOne := convertActiveNumberToStr + 1
		convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
		poNumber := getFrontDigitNumber + convertAddedActiveNumberToStr
		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		finalNumber = poNumber
		if p.Type == "proforma" {
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "PRF", monthInRoman, year)
		}
		if p.Type == "invoice" {
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "INV", monthInRoman, year)
		}
	}

	Invoice := models.Invoice{
		Number:              formattedNumber,
		Type:                invoiceType,
		Date:                p.Date,
		Sequence:            finalNumber,
		OperatingActivityID: p.OperatingActivityID,
		InvoiceCreatedBy:    p.CreatedBy,
	}
	if InvoiceErr := tx.Create(&Invoice); err != nil {
		tx.Rollback()
		return InvoiceErr.Error
	}
	tx.Commit()
	return
}

func (module *module) UpdateInvoice(id uuid.UUID, p *dto.UpdateInvoice) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	Invoice := models.Invoice{
		ID:                  id,
		Type:                p.Type,
		Date:                p.Date,
		OperatingActivityID: p.OperatingActivityID,
		InvoiceUpdatedBy:    p.UpdatedBy,
	}
	if InvoiceErr := tx.Updates(&Invoice); err != nil {
		tx.Rollback()
		return InvoiceErr.Error
	}
	tx.Commit()
	return
}

func (module *module) CheckExistingInvoice(id string, param CheckExistingInvoiceStruct) (err error) {

	if param.Number != "" {
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByOperatingActivityID(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
