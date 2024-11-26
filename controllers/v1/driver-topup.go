package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/services"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

// AuthLogin godoc
// @Summary      Method to get all driver topup by drivers
// @Description  A GET Request to fetch a all records driver topup
// @Tags         Driver Topup
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /driver-topup/get [get]
func GETAllDriverTopup(c *gin.Context) {
	if driverTopup, err := services.Handler.RetrieveAllDriverTopup(); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver topup"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &driverTopup, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to insert new row of driver topups
// @Description  A POST Request to insert new row of driver topups
// @Tags         Driver Topup
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data formData dto.InsertFormDriverTopup true "insert driver topup"
// @Param receiptImage formData file true "receipt image upload"
// @Router       /driver-topup/create [post]
func POSTCreateDriverTopup(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.InsertFormDriverTopup{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fReceiptImageDocument, _ := c.FormFile("receiptImage")
	if fReceiptImageDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field receiptImage is required"), "cannot submit if receipt image is empty"))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver"))
		return
	}

	topupAmount, _ := strconv.ParseFloat(pValidator.Amount, 2)

	p := &dto.InsertDriverTopup{
		Amount:       topupAmount,
		ReceiptImage: fReceiptImageDocument,
		DriverID:     cred.CredentialDriver.ID,
	}

	if err = services.Handler.InsertDriverTopup(c, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "driver topup"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
