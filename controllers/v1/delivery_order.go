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
		return
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
	return
}

func POSTDeliveryOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	pValidator := &dto.InsertDeliveryOrderValidator{CreatedBy: userId}
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
		ContactPerson:       pValidator.ContactPerson,
		PhoneNumber:         pValidator.PhoneNumber,
		Address:             pValidator.Address,
		Note:                pValidator.Note,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		Status:              pValidator.Status,
		OperatingActivityID: operatingActivityId,
		CreatedBy:           pValidator.CreatedBy,
	}

	// check existing operating id
	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err = services.Handler.InsertDeliveryOrder(c, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "delivery order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
	return
}

func PUTDeliveryOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	pValidator := &dto.UpdateDeliveryOrderValidator{UpdatedBy: userId}
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
		ContactPerson:       pValidator.ContactPerson,
		PhoneNumber:         pValidator.PhoneNumber,
		Address:             pValidator.Address,
		Note:                pValidator.Note,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		Status:              pValidator.Status,
		OperatingActivityID: operatingActivityId,
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
	return
}
