package v1

import (
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
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

	var operatingActivity []models.OperatingActivity
	if operatingActivity, err = services.Handler.RetrieveAllOperatingActivity(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: operatingActivity})
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
}

func POSTOperatingActivity(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.InsertOperatingActivity{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
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
}

func PUTOperatingActivity(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.UpdateOperatingActivity{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	operatingActivityId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if operatingActivity, err := services.Handler.RetrieveOperatingActivity(operatingActivityId); err == nil {

		if operatingActivity.TaxInvoiceNumber != p.TaxInvoiceNumber {
			if err := services.Handler.CheckExistingOperatingActivity(id, struct{ *models.OperatingActivity }{&models.OperatingActivity{
				TaxInvoiceNumber: p.TaxInvoiceNumber,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, p.TaxInvoiceNumber))
				return
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity"))
		return
	}

	if err = services.Handler.UpdateOperatingActivity(operatingActivityId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "operating activity"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
