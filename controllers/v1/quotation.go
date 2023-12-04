package v1

import (
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

func GETAllQuotation(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if quotation, err := services.Handler.RetrieveAllQuotationPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: quotation})
	}

}

func GETQuotation(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	quotationId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var quotation models.Quotation
	if quotation, err = services.Handler.RetrieveQuotation(quotationId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: quotation})
}

func POSTQuotation(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertQuotationValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)

	p := &dto.InsertQuotation{
		Number:              pValidator.Number,
		FrancoArea:          pValidator.FrancoArea,
		PaymentTerm:         pValidator.PaymentTerm,
		SendAfter:           pValidator.SendAfter,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		OperatingActivityID: operatingActivityId,
		CreatedBy:           pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingQuotation("", struct{ *models.Quotation }{&models.Quotation{
		Number: p.Number,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, ""))
		return
	}

	if err = services.Handler.InsertQuotation(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "quotation"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTQuotation(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateQuotationValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.UpdateQuotation{
		Number:              pValidator.Number,
		FrancoArea:          pValidator.FrancoArea,
		PaymentTerm:         pValidator.PaymentTerm,
		SendAfter:           pValidator.SendAfter,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		OperatingActivityID: operatingActivityId,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	quotationId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if quotation, err := services.Handler.RetrieveQuotation(quotationId); err == nil {

		if quotation.Number != p.Number {
			if err := services.Handler.CheckExistingQuotation(id, struct{ *models.Quotation }{&models.Quotation{
				Number: p.Number,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, p.Number))
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	}

	if err = services.Handler.UpdateQuotation(quotationId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "quotation"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
