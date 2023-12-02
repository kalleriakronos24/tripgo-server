package v1

import (
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllPayment(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var Payment []models.Payment
	if Payment, err = services.Handler.RetrieveAllPayment(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: Payment})
}

func GETPayment(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	PaymentId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var Payment models.Payment
	if Payment, err = services.Handler.RetrievePayment(PaymentId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: Payment})
}

func POSTPayment(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertPaymentValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.InsertPayment{
		Currency:            pValidator.Currency,
		Amount:              pValidator.Amount,
		Term:                pValidator.Term,
		ARAging:             pValidator.ARAging,
		Remarks:             pValidator.Remarks,
		DueDate:             utils.ConvertStrToDateTime(pValidator.DueDate),
		SettlementDate:      utils.ConvertStrToDateTime(pValidator.SettlementDate),
		OperatingActivityID: operatingActivityId,
		CreatedBy:           pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err = services.Handler.InsertPayment(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "payment"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTPayment(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdatePaymentValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.UpdatePayment{
		Currency:            pValidator.Currency,
		Amount:              pValidator.Amount,
		Term:                pValidator.Term,
		ARAging:             pValidator.ARAging,
		Remarks:             pValidator.Remarks,
		DueDate:             utils.ConvertStrToDateTime(pValidator.DueDate),
		SettlementDate:      utils.ConvertStrToDateTime(pValidator.SettlementDate),
		OperatingActivityID: operatingActivityId,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	PaymentId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if _, err := services.Handler.RetrievePayment(PaymentId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment"))
		return
	}

	if err = services.Handler.UpdatePayment(PaymentId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "payment"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
