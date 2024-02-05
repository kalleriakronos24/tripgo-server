package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	"github.com/skip2/go-qrcode"
	"gitlab.com/odma1/odma-be/config"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"path"
	"strconv"
	"strings"
	"time"
)

type CheckExistingPurchaseOrderStruct struct {
	*models.PurchaseOrder
}

func (module *module) RetrieveAllPurchaseOrder(id uuid.UUID) (m []models.PurchaseOrder, err error) {
	if m, err = module.db.purchaseOrderModel.GetAllPurchaseOrder(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllPurchaseOrderPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.purchaseOrderModel.GetAllPurchaseOrderPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrievePurchaseOrder(id uuid.UUID) (m models.CustomResponsePurchaseOrder, err error) {
	if m, err = module.db.purchaseOrderModel.GetOnePurchaseOrderByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertPurchaseOrder(c *gin.Context, p *dto.InsertPurchaseOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var operatingActivity models.OperatingActivity
	var operatingActivityErr error

	if operatingActivity, operatingActivityErr = module.db.operatingActivityModel.GetOneOperatingActivityByID(p.OperatingActivityID); operatingActivityErr != nil {
		return errors.New(operatingActivityErr.Error())
	}

	if p.Type == "in" {
		clientData := operatingActivity.Client
		documentPath := fmt.Sprintf("/client/%s/purchase-order/%s", clientData.Name, p.Document.Filename)
		fileExt := path.Ext(p.Document.Filename)
		document := models.Document{
			Base64:            "",
			Path:              documentPath,
			AbsolutePath:      fmt.Sprintf("%s/%s/%s/client/%s/purchase-order/%s", config.AppConfig.APPUrl, config.AppConfig.APPUrlStaticFileGroupRoute, config.AppConfig.AppUrlStaticFileMainRoute, clientData.Name, p.Document.Filename),
			FileName:          p.Document.Filename,
			Extension:         fileExt,
			Location:          "local",
			DocumentCreatedBy: p.CreatedBy,
		}
		if errDocument := tx.Create(&document); errDocument.Error != nil {
			tx.Rollback()
			return errDocument.Error
		}
		documentTypePath := fmt.Sprintf("%s/purchase-order", clientData.Name)
		if _, saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, nil, nil); saveFileErr != nil {
			return errors.New(saveFileErr.Error())
		}
		purchaseOrder := models.PurchaseOrder{
			Number:                 strings.ToUpper(p.Number),
			Type:                   p.Type,
			Recipient:              p.Recipient,
			RecipientEmail:         p.RecipientEmail,
			Date:                   p.Date,
			OperatingActivityID:    p.OperatingActivityID,
			DocumentID:             document.ID,
			PurchaseOrderCreatedBy: p.CreatedBy,
		}

		if purchaseOrderErr := tx.Create(&purchaseOrder); purchaseOrderErr.Error != nil {
			tx.Rollback()
			return purchaseOrderErr.Error
		}
	}

	if p.Type == "out" {
		purchaseOrder := models.PurchaseOrder{
			Number:                 strings.ToUpper(p.Number),
			Type:                   p.Type,
			Recipient:              p.Recipient,
			RecipientEmail:         p.RecipientEmail,
			Date:                   p.Date,
			OperatingActivityID:    p.OperatingActivityID,
			PurchaseOrderCreatedBy: p.CreatedBy,
		}

		if purchaseOrderErr := tx.Create(&purchaseOrder); purchaseOrderErr.Error != nil {
			tx.Rollback()
			return purchaseOrderErr.Error
		}
	}

	tx.Commit()
	return
}

func (module *module) UpdatePurchaseOrder(c *gin.Context, id uuid.UUID, p *dto.UpdatePurchaseOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	if p.Type == "in" {

		var operatingActivity models.OperatingActivity
		var operatingActivityErr error

		if operatingActivity, operatingActivityErr = module.db.operatingActivityModel.GetOneOperatingActivityByID(p.OperatingActivityID); operatingActivityErr != nil {
			return errors.New(operatingActivityErr.Error())
		}

		currentPurchaseOrder, err := module.RetrievePurchaseOrder(id)

		if err != nil {
			return errors.New(err.Error())
		}

		currentDocument, err := module.RetrieveDocument(currentPurchaseOrder.DocumentID)

		if err != nil {
			return errors.New(err.Error())
		}

		/**
		in here we just update the document's values regardless the user upload a new file or the same

		because the docker logic, if it's same file, it will replaces the old one.

		if it's different than the old one, it will adds the new file, keeping the old one.

		but since we update the values in our table, we will retrieve the latest file that user updated
		*/
		clientData := operatingActivity.Client
		documentPath := fmt.Sprintf("/client/%s/purchase-order/%s", clientData.Name, p.Document.Filename)
		fileExt := path.Ext(p.Document.Filename)

		document := models.Document{
			ID:                currentDocument.ID,
			Base64:            "",
			Path:              documentPath,
			AbsolutePath:      fmt.Sprintf("%s/%s/%s/client/%s/purchase-order/%s", config.AppConfig.APPUrl, config.AppConfig.APPUrlStaticFileGroupRoute, config.AppConfig.AppUrlStaticFileMainRoute, clientData.Name, p.Document.Filename),
			FileName:          p.Document.Filename,
			Extension:         fileExt,
			Location:          "local",
			DocumentUpdatedBy: p.UpdatedBy,
		}

		if errDocument := tx.Updates(&document); errDocument.Error != nil {
			tx.Rollback()
			return errDocument.Error
		}

		documentTypePath := fmt.Sprintf("%s/purchase-order", clientData.Name)

		if _, saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, p.Document, nil); saveFileErr != nil {
			return errors.New(saveFileErr.Error())
		}

		purchaseOrder := models.PurchaseOrder{
			ID:                     id,
			Number:                 strings.ToUpper(p.Number),
			Type:                   p.Type,
			Recipient:              p.Recipient,
			RecipientEmail:         p.RecipientEmail,
			Date:                   p.Date,
			OperatingActivityID:    p.OperatingActivityID,
			PurchaseOrderUpdatedBy: p.UpdatedBy,
		}

		if purchaseOrderErr := tx.Updates(&purchaseOrder); purchaseOrderErr.Error != nil {
			tx.Rollback()
			return purchaseOrderErr.Error
		}
	}

	if p.Type == "out" {

		PurchaseOrderModel := models.PurchaseOrder{}

		_, purchaseOrderLastDataErr := database.GetLastDocumentNumber(tx, &PurchaseOrderModel)

		var formattedNumber string
		var finalNumber string
		if purchaseOrderLastDataErr != nil {
			now := time.Now()
			year := now.Year()
			month := now.Month()
			monthInRoman := utils.IntegerToRoman(int(month))
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", "0001", "PO", monthInRoman, year)
			finalNumber = "0001"
		} else {
			parsedNumber, _ := strconv.Atoi(PurchaseOrderModel.Sequence)
			strLastNumber := strconv.Itoa(parsedNumber)
			getFrontDigitNumber := strings.Replace(PurchaseOrderModel.Sequence, strLastNumber, "", -1)
			convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
			addActiveNumberByOne := convertActiveNumberToStr + 1
			convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
			poNumber := getFrontDigitNumber + convertAddedActiveNumberToStr

			now := time.Now()
			year := now.Year()
			month := now.Month()
			monthInRoman := utils.IntegerToRoman(int(month))
			formattedNumber = fmt.Sprintf("%s/%s/%s/%d", poNumber, "PO", monthInRoman, year)
			finalNumber = poNumber
		}

		purchaseOrder := models.PurchaseOrder{
			ID:                     id,
			Number:                 formattedNumber,
			Type:                   p.Type,
			Recipient:              p.Recipient,
			RecipientEmail:         p.RecipientEmail,
			Date:                   p.Date,
			Sequence:               finalNumber,
			OperatingActivityID:    p.OperatingActivityID,
			PurchaseOrderUpdatedBy: p.UpdatedBy,
		}

		if purchaseOrderErr := tx.Updates(&purchaseOrder); purchaseOrderErr.Error != nil {
			tx.Rollback()
			return purchaseOrderErr.Error
		}
	}

	tx.Commit()
	return
}

type RemappedProduct struct {
	Quantity   float64
	VATRate    int16
	SubTotal   string
	GrandTotal string
	Product    models.CustomProduct
}

func (module *module) GeneratePurchaseOrderDocument(id uuid.UUID) (output GenerateDocumentOutput, err error) {

	formatCurrency := accounting.Accounting{Symbol: "", Precision: 2}

	var purchaseOrder models.CustomResponsePurchaseOrder
	if purchaseOrder, err = module.RetrievePurchaseOrder(id); err != nil {
		return output, errors.New("purchase order not found")
	}

	var subTotalProduct float64
	var grandTotalProduct string
	var ppnProduct string

	var remappedProduct []RemappedProduct

	if len(purchaseOrder.Product) > 0 {
		for _, data := range purchaseOrder.Product {
			var x models.PurchaseOrderProduct = data.(models.PurchaseOrderProduct)
			subTotalProduct += x.SubTotal

			subTotalProductFormatted := formatCurrency.FormatMoney(x.SubTotal)
			unitPriceProductFormatted := formatCurrency.FormatMoney(x.Product.UnitPrice * 0.11)
			grandTotalProduct = formatCurrency.FormatMoney(x.SubTotal + (x.SubTotal * 0.11))

			productStruct := models.CustomProduct{
				Name:      x.Product.Name,
				UnitPrice: unitPriceProductFormatted,
				Packaging: x.Product.Packaging,
				Stock:     x.Product.Stock,
				Note:      x.Product.Note,
			}
			productData := RemappedProduct{
				Quantity:   x.Quantity,
				VATRate:    x.VATRate,
				SubTotal:   subTotalProductFormatted,
				GrandTotal: grandTotalProduct,
				Product:    productStruct,
			}
			remappedProduct = append(remappedProduct, productData)
		}
	} else {
		return output, errors.New("cannot generate if there is no product in this purchase order")
	}

	subTotalProductFormatted := formatCurrency.FormatMoney(subTotalProduct)
	ppnProduct = formatCurrency.FormatMoney(subTotalProduct * 0.11)
	grandTotalProduct = formatCurrency.FormatMoney(subTotalProduct + (subTotalProduct * 0.11))

	date := utils.ConvertEnToIDDateTime(purchaseOrder.Date)

	png, err := qrcode.Encode(purchaseOrder.Number, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	dataURI := base64.StdEncoding.EncodeToString([]byte(png))

	templateData := struct {
		ClientName      string
		ClientAddress   string
		Recipient       string
		RecipientEmails string
		DocumentNumber  string
		DateGenerated   string
		Products        []RemappedProduct
		SubTotal        string
		Ppn             string
		GrandTotal      string
		QrCode          string
	}{
		ClientName:      purchaseOrder.OperatingActivity.Client.Name,
		ClientAddress:   purchaseOrder.OperatingActivity.Client.Address,
		Recipient:       purchaseOrder.Recipient,
		RecipientEmails: purchaseOrder.RecipientEmail,
		DocumentNumber:  purchaseOrder.Number,
		DateGenerated:   date,
		Products:        remappedProduct,
		SubTotal:        subTotalProductFormatted,
		Ppn:             ppnProduct,
		GrandTotal:      grandTotalProduct,
		QrCode:          dataURI,
	}

	var outputPath string
	var saveFileErr error
	if outputPath, saveFileErr = utils.SaveFileToDockerVolume(nil, "", "po-out", nil, templateData); saveFileErr != nil {
		return output, errors.New(saveFileErr.Error())
	}

	data := GenerateDocumentOutput{
		OutputPath: outputPath,
		FileName:   purchaseOrder.Number,
	}

	return data, nil
}

func (module *module) CheckExistingPurchaseOrder(id string, param CheckExistingPurchaseOrderStruct) (err error) {

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.purchaseOrderModel.GetOnePurchaseOrderByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}

	if param.OperatingActivityID.String() != "00000000-0000-0000-0000-000000000000" {
		_, optActIdErr := uuid.Parse(param.OperatingActivityID.String())
		if optActIdErr == nil {
			if _, dbErr := module.db.purchaseOrderModel.GetOnePurchaseOrderByOperatingActivityId(param.OperatingActivityID); dbErr != nil {
				return errors.New(dbErr.Error())
			}
			return
		}
		return
	}

	return
}
