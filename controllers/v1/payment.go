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
	p := &dto.InsertPayment{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
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
	p := &dto.UpdatePayment{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	PaymentId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
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
