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

func GETAllPaymentInstallment(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PaymentInstallment []models.PaymentInstallment
	if PaymentInstallment, err = services.Handler.RetrieveAllPaymentInstallment(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment installment"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: PaymentInstallment})
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
	p := &dto.InsertPaymentInstallment{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
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
	p := &dto.UpdatePaymentInstallment{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	PaymentInstallmentId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
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
