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
	"log"
	"path"
)

type CheckExistingDeliveryOrderStruct struct {
	*models.DeliveryOrder
}

func (module *module) RetrieveDeliveryOrder(id uuid.UUID) (m models.DeliveryOrder, err error) {
	if m, err = module.db.deliveryOrderModel.GetOneDeliveryOrderByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllDeliveryOrder(id uuid.UUID) (m []models.DeliveryOrder, err error) {
	if m, err = module.db.deliveryOrderModel.GetAllDeliveryOrder(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertDeliveryOrder(c *gin.Context, p *dto.InsertDeliveryOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var operatingActivity models.OperatingActivity
	var operatingActivityErr error

	if operatingActivity, operatingActivityErr = module.db.operatingActivityModel.GetOneOperatingActivityByID(p.OperatingActivityID); operatingActivityErr != nil {
		return errors.New(operatingActivityErr.Error())
	}
	log.Println("DDDD >>> ", p.Number)
	clientData := operatingActivity.Client
	documentPath := fmt.Sprintf("/client/%s/delivery-order/%s", clientData.Name, p.Document.Filename)
	fileExt := path.Ext(p.Document.Filename)

	document := models.Document{
		Base64:            "",
		Path:              documentPath,
		AbsolutePath:      fmt.Sprintf("%s/%s/%s/client/%s/delivery-order/%s", config.AppConfig.APPUrl, config.AppConfig.APPUrlStaticFileGroupRoute, config.AppConfig.AppUrlStaticFileMainRoute, clientData.Name, p.Document.Filename),
		FileName:          p.Document.Filename,
		Extension:         fileExt,
		Location:          "local",
		DocumentCreatedBy: p.CreatedBy,
	}
	if errDocument := tx.Create(&document); err != nil {
		tx.Rollback()
		return errDocument.Error
	}

	documentTypePath := fmt.Sprintf("%s/delivery-order", clientData.Name)

	if saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, p.Document); saveFileErr != nil {
		return errors.New(saveFileErr.Error())
	}

	DeliveryOrder := models.DeliveryOrder{
		Number:                 p.Number,
		ContactPerson:          p.ContactPerson,
		PhoneNumber:            p.PhoneNumber,
		Address:                p.Address,
		Note:                   p.Note,
		Date:                   p.Date,
		Status:                 p.Status,
		DocumentID:             document.ID,
		OperatingActivityID:    p.OperatingActivityID,
		DeliveryOrderCreatedBy: p.CreatedBy,
	}
	if DeliveryOrderErr := tx.Create(&DeliveryOrder); err != nil {
		tx.Rollback()
		return DeliveryOrderErr.Error
	}
	tx.Commit()
	return
}

func (module *module) UpdateDeliveryOrder(c *gin.Context, id uuid.UUID, p *dto.UpdateDeliveryOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var operatingActivity models.OperatingActivity
	var operatingActivityErr error

	if operatingActivity, operatingActivityErr = module.db.operatingActivityModel.GetOneOperatingActivityByID(p.OperatingActivityID); operatingActivityErr != nil {
		return errors.New(operatingActivityErr.Error())
	}

	currentDeliveryOrder, err := module.RetrieveDeliveryOrder(id)

	if err != nil {
		return errors.New(err.Error())
	}

	currentDocument, err := module.RetrieveDocument(currentDeliveryOrder.DocumentID)

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
	documentPath := fmt.Sprintf("/client/%s/delivery-order/%s", clientData.Name, p.Document.Filename)
	fileExt := path.Ext(p.Document.Filename)

	document := models.Document{
		ID:                currentDocument.ID,
		Base64:            "",
		Path:              documentPath,
		AbsolutePath:      fmt.Sprintf("%s/%s/%s/client/%s/delivery-order/%s", config.AppConfig.APPUrl, config.AppConfig.APPUrlStaticFileGroupRoute, config.AppConfig.AppUrlStaticFileMainRoute, clientData.Name, p.Document.Filename),
		FileName:          p.Document.Filename,
		Extension:         fileExt,
		Location:          "local",
		DocumentUpdatedBy: p.UpdatedBy,
	}

	if errDocument := tx.Updates(&document); err != nil {
		tx.Rollback()
		return errDocument.Error
	}

	documentTypePath := fmt.Sprintf("%s/delivery-order", clientData.Name)

	if saveFileErr := utils.SaveFileToDockerVolume(c, "client", documentTypePath, p.Document); saveFileErr != nil {
		return errors.New(saveFileErr.Error())
	}

	DeliveryOrder := models.DeliveryOrder{
		ID:                     id,
		Number:                 p.Number,
		ContactPerson:          p.ContactPerson,
		PhoneNumber:            p.PhoneNumber,
		Address:                p.Address,
		Note:                   p.Note,
		Date:                   p.Date,
		Status:                 p.Status,
		DocumentID:             document.ID,
		OperatingActivityID:    p.OperatingActivityID,
		DeliveryOrderUpdatedBy: p.UpdatedBy,
	}
	if DeliveryOrderErr := tx.Updates(&DeliveryOrder); err != nil {
		tx.Rollback()
		return DeliveryOrderErr.Error
	}
	tx.Commit()
	return
}

func (module *module) CheckExistingDeliveryOrder(id string, param CheckExistingDeliveryOrderStruct) (err error) {

	if param.Number != "" {
		log.Println("DO NUMBER ", param.Number)
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByOperatingActivityID(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
