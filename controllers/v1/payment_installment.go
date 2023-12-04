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

func GETAllPaymentInstallment(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if PaymentInstallment, err := services.Handler.RetrieveAllPaymentInstallmentPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment installment"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: PaymentInstallment})
	}

}

func GETPaymentInstallment(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	PaymentInstallmentId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PaymentInstallment models.PaymentInstallment
	if PaymentInstallment, err = services.Handler.RetrievePaymentInstallment(PaymentInstallmentId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment installment"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: PaymentInstallment})
}

func POSTPaymentInstallment(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertPaymentInstallmentValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	paymentId, _ := uuid.Parse(pValidator.PaymentID)
	p := &dto.InsertPaymentInstallment{
		Currency:  pValidator.Currency,
		Amount:    pValidator.Amount,
		Note:      pValidator.Note,
		Date:      utils.ConvertStrToDateTime(pValidator.Date),
		PaymentID: paymentId,
		CreatedBy: pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingPayment(p.PaymentID.String(), struct{ *models.Payment }{&models.Payment{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("payment id %s is not found", p.PaymentID)))
		return
	}

	if err = services.Handler.InsertPaymentInstallment(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "payment installment"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTPaymentInstallment(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdatePaymentInstallmentValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	paymentId, _ := uuid.Parse(pValidator.PaymentID)
	p := &dto.UpdatePaymentInstallment{
		Currency:  pValidator.Currency,
		Amount:    pValidator.Amount,
		Note:      pValidator.Note,
		Date:      utils.ConvertStrToDateTime(pValidator.Date),
		PaymentID: paymentId,
		UpdatedBy: pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	PaymentInstallmentId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingPayment(p.PaymentID.String(), struct{ *models.Payment }{&models.Payment{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("payment id %s is not found", p.PaymentID)))
		return
	}

	if _, err := services.Handler.RetrievePaymentInstallment(PaymentInstallmentId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment installment"))
		return
	}

	if err = services.Handler.UpdatePaymentInstallment(PaymentInstallmentId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "payment installment"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
