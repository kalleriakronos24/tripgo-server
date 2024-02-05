package v1

import (
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllOperatingActivity(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if operatingActivity, err := services.Handler.RetrieveAllOperatingActivityPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: operatingActivity})
		return
	}
}

func GETOperatingActivity(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	operatingActivityId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var operatingActivity models.OperatingActivity
	if operatingActivity, err = services.Handler.RetrieveOperatingActivity(operatingActivityId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: operatingActivity})
	return
}

func POSTOperatingActivity(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	pValidator := &dto.InsertOperatingActivityValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	clientId, _ := uuid.Parse(pValidator.ClientID)
	p := &dto.InsertOperatingActivity{
		TaxInvoiceNumber:      pValidator.TaxInvoiceNumber,
		DeliveryReceiptNumber: pValidator.DeliveryReceiptNumber,
		Status:                pValidator.Status,
		ClientID:              clientId,
		CreatedBy:             pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingClient(p.ClientID.String(), struct{ *masterModels.Client }{&masterModels.Client{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("client id %s is not found", p.ClientID)))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity("", struct{ *models.OperatingActivity }{&models.OperatingActivity{
		TaxInvoiceNumber: p.TaxInvoiceNumber,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, ""))
		return
	}

	if err = services.Handler.InsertOperatingActivity(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "operating activity"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
	return
}

func PUTOperatingActivity(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateOperatingActivityValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	clientId, _ := uuid.Parse(pValidator.ClientID)
	p := &dto.UpdateOperatingActivity{
		ID:                    pValidator.ID,
		TaxInvoiceNumber:      pValidator.TaxInvoiceNumber,
		DeliveryReceiptNumber: pValidator.DeliveryReceiptNumber,
		Status:                pValidator.Status,
		ClientID:              clientId,
		UpdatedBy:             pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	operatingActivityId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingClient(p.ClientID.String(), struct{ *masterModels.Client }{&masterModels.Client{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("client id %s is not found", p.ClientID)))
		return
	}

	if operatingActivity, err := services.Handler.RetrieveOperatingActivity(operatingActivityId); err == nil {

		if operatingActivity.TaxInvoiceNumber != p.TaxInvoiceNumber {
			if err := services.Handler.CheckExistingOperatingActivity(id, struct{ *models.OperatingActivity }{&models.OperatingActivity{
				TaxInvoiceNumber: p.TaxInvoiceNumber,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, p.TaxInvoiceNumber))
				return
			} else {
				// runs the update query
				if err = services.Handler.UpdateOperatingActivity(operatingActivityId, p); err != nil {
					c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "operating activity"))
					return
				} else {
					c.JSON(http.StatusOK, dto.Response{Message: "success"})
					return
				}
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity"))
		return
	}

	if err = services.Handler.UpdateOperatingActivity(operatingActivityId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "operating activity"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Message: "success"})
		return
	}
}
