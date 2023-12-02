package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/services"
	"log"
	"net/http"
)

func GETAllDeliveryOrder(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var DeliveryOrder []models.DeliveryOrder
	if DeliveryOrder, err = services.Handler.RetrieveAllDeliveryOrder(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "delivery order"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: DeliveryOrder})
}

func GETDeliveryOrder(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	DeliveryOrderId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var DeliveryOrder models.DeliveryOrder
	if DeliveryOrder, err = services.Handler.RetrieveDeliveryOrder(DeliveryOrderId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "delivery order"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: DeliveryOrder})
}

func POSTDeliveryOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	fDeliveryOrderDocument, _ := c.FormFile("document")

	if fDeliveryOrderDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field document is required"), "cannot submit if document is empty"))
		return
	}

	p := &dto.InsertFormDataDeliveryOrder{CreatedBy: userId}

	//var data *multipart.Form
	if err = c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityID, err := uuid.Parse(p.OperatingActivityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "invalid operating activity id"))
		return
	}

	repopulateFormDataPayload := &dto.InsertDeliveryOrder{
		Number:              p.Number,
		Date:                p.Date,
		Status:              p.Status,
		Note:                p.Note,
		Address:             p.Address,
		PhoneNumber:         p.PhoneNumber,
		ContactPerson:       p.ContactPerson,
		Document:            fDeliveryOrderDocument,
		OperatingActivityID: operatingActivityID,
		CreatedBy:           p.CreatedBy,
	}

	// check existing operating id
	if err := services.Handler.CheckExistingOperatingActivity(operatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	// check duplication delivery number
	if err := services.Handler.CheckExistingDeliveryOrder("", struct {
		*models.DeliveryOrder
	}{&models.DeliveryOrder{
		Number: repopulateFormDataPayload.Number,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with delivery number %s", p.Number)))
		return
	}

	if err = services.Handler.InsertDeliveryOrder(c, repopulateFormDataPayload); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "delivery order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTDeliveryOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	fDeliveryOrderDocument, _ := c.FormFile("document")

	if fDeliveryOrderDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field document is required"), "cannot submit if document is empty"))
		return
	}
	p := &dto.UpdateFormDataDeliveryOrder{UpdatedBy: userId}

	if err = c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	id, _ := c.Params.Get("id")
	DeliveryOrderId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	operatingActivityID, err := uuid.Parse(p.OperatingActivityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "invalid operating activity id"))
		return
	}

	repopulateFormDataPayload := &dto.UpdateDeliveryOrder{
		Number:              p.Number,
		Date:                p.Date,
		Status:              p.Status,
		Note:                p.Note,
		Address:             p.Address,
		PhoneNumber:         p.PhoneNumber,
		ContactPerson:       p.ContactPerson,
		Document:            fDeliveryOrderDocument,
		OperatingActivityID: operatingActivityID,
	}

	if DeliveryOrder, err := services.Handler.RetrieveDeliveryOrder(DeliveryOrderId); err == nil {

		// check duplication delivery number
		if DeliveryOrder.Number != repopulateFormDataPayload.Number {
			if err := services.Handler.CheckExistingDeliveryOrder(id, struct {
				*models.DeliveryOrder
			}{&models.DeliveryOrder{
				Number: repopulateFormDataPayload.Number,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with delivery number %s", DeliveryOrder.Number)))
				return
			}
		}

		// check existing operating id
		if err := services.Handler.CheckExistingOperatingActivity(operatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
			return
		}

		// check duplication operating id
		if DeliveryOrder.OperatingActivityID != repopulateFormDataPayload.OperatingActivityID {
			if err := services.Handler.CheckExistingDeliveryOrder(id, struct {
				*models.DeliveryOrder
			}{&models.DeliveryOrder{
				OperatingActivityID: repopulateFormDataPayload.OperatingActivityID,
			}}); err != nil {

				log.Println(err)
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", DeliveryOrder.OperatingActivityID)))
				return
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "delivery order"))
		return
	}

	if err = services.Handler.UpdateDeliveryOrder(c, DeliveryOrderId, repopulateFormDataPayload); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "delivery order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
