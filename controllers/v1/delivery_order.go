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
	"gitlab.com/odma1/odma-be/utils"
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

	if DeliveryOrder, err := services.Handler.RetrieveAllDeliveryOrderPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "delivery order"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: DeliveryOrder})
	}
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

	pValidator := &dto.InsertFormDataDeliveryOrder{CreatedBy: userId}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.InsertDeliveryOrder{
		Number:              pValidator.Number,
		ContactPerson:       pValidator.ContactPerson,
		PhoneNumber:         pValidator.PhoneNumber,
		Address:             pValidator.Address,
		Note:                pValidator.Note,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		Status:              pValidator.Status,
		OperatingActivityID: operatingActivityId,
		Document:            fDeliveryOrderDocument,
		CreatedBy:           pValidator.CreatedBy,
	}

	// check existing operating id
	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	// check duplication delivery number
	if err := services.Handler.CheckExistingDeliveryOrder("", struct {
		*models.DeliveryOrder
	}{&models.DeliveryOrder{
		Number: p.Number,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with delivery number %s", p.Number)))
		return
	}

	if err = services.Handler.InsertDeliveryOrder(c, p); err != nil {
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
	pValidator := &dto.UpdateFormDataDeliveryOrder{UpdatedBy: userId}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.UpdateDeliveryOrder{
		ID:                  pValidator.ID,
		Number:              pValidator.Number,
		ContactPerson:       pValidator.ContactPerson,
		PhoneNumber:         pValidator.PhoneNumber,
		Address:             pValidator.Address,
		Note:                pValidator.Note,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		Status:              pValidator.Status,
		OperatingActivityID: operatingActivityId,
		Document:            fDeliveryOrderDocument,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	DeliveryOrderId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("invalid uuid format"), "invalid uuid delivery order id parameter"))
		return
	}

	// checks
	if DeliveryOrder, err := services.Handler.RetrieveDeliveryOrder(DeliveryOrderId); err == nil {

		// check duplication delivery number
		if DeliveryOrder.Number != p.Number {
			if err := services.Handler.CheckExistingDeliveryOrder(id, struct {
				*models.DeliveryOrder
			}{&models.DeliveryOrder{
				Number: p.Number,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with delivery number %s", DeliveryOrder.Number)))
				return
			}
		}

		// check existing operating id
		if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
			return
		}

		// check duplication operating id
		if DeliveryOrder.OperatingActivityID != p.OperatingActivityID {
			if err := services.Handler.CheckExistingDeliveryOrder(id, struct {
				*models.DeliveryOrder
			}{&models.DeliveryOrder{
				OperatingActivityID: p.OperatingActivityID,
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

	if err = services.Handler.UpdateDeliveryOrder(c, DeliveryOrderId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "delivery order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
