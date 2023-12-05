package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/config"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"path"
	"strings"
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
	if errDocument := tx.Create(&document); err != nil {
		tx.Rollback()
		return errDocument.Error
	}

	documentTypePath := fmt.Sprintf("%s/purchase-order", clientData.Name)

	if saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, p.Document); saveFileErr != nil {
		return errors.New(saveFileErr.Error())
	}

	purchaseOrder := models.PurchaseOrder{
		Number:                 strings.ToUpper(p.Number),
		Type:                   p.Type,
		Recipient:              p.Recipient,
		RecipientEmail:         p.RecipientEmail,
		Date:                   p.Date,
		DocumentID:             document.ID,
		OperatingActivityID:    p.OperatingActivityID,
		PurchaseOrderCreatedBy: p.CreatedBy,
	}
	if purchaseOrderErr := tx.Create(&purchaseOrder); err != nil {
		tx.Rollback()
		return purchaseOrderErr.Error
	}
	tx.Commit()
	return
}

func (module *module) UpdatePurchaseOrder(c *gin.Context, id uuid.UUID, p *dto.UpdatePurchaseOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

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

	if errDocument := tx.Updates(&document); err != nil {
		tx.Rollback()
		return errDocument.Error
	}

	documentTypePath := fmt.Sprintf("%s/purchase-order", clientData.Name)

	if saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, p.Document); saveFileErr != nil {
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
	if purchaseOrderErr := tx.Updates(&purchaseOrder); err != nil {
		tx.Rollback()
		return purchaseOrderErr.Error
	}
	tx.Commit()
	return
}

func (module *module) CheckExistingPurchaseOrder(id string, param CheckExistingPurchaseOrderStruct) (err error) {

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.purchaseOrderModel.GetOnePurchaseOrderByOperatingActivityId(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.purchaseOrderModel.GetOnePurchaseOrderByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
