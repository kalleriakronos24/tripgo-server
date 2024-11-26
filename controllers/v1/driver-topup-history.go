package v1

import (
	"net/http"

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
// @Tags         Driver Topup History
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /driver-topup-history/get [get]
func GETAllDriverTopupByDriver(c *gin.Context) {

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver"))
		return
	}

	if driverTopup, err := services.Handler.RetrieveAllDriverTopupByDriver(cred.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "topup history"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &driverTopup, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to approve driver topup request
// @Description  A POST Request to approve driver topup request
// @Tags         Driver Topup History
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /driver-topup-history/approve [post]
func POSTApproveDriverTopup(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.UpdateDriverTopupHistory{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver"))
		return
	}

	p := &dto.UpdateDriverTopupHistory{
		Amount:        pValidator.Amount,
		DriverID:      cred.CredentialDriver.ID,
		DriverTopupID: pValidator.DriverTopupID,
		Status:        pValidator.Status,
	}

	if err = services.Handler.ApproveDriverTopup(c, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "driver topup"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to reject driver topup request
// @Description  A POST Request to reject driver topup request
// @Tags         Driver Topup History
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /driver-topup-history/reject [post]
func POSTRejectDriverTopup(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.UpdateDriverTopupHistory{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver"))
		return
	}

	p := &dto.UpdateDriverTopupHistory{
		Amount:        pValidator.Amount,
		DriverID:      cred.CredentialDriver.ID,
		DriverTopupID: pValidator.DriverTopupID,
		Status:        pValidator.Status,
	}

	if err = services.Handler.ApproveDriverTopup(c, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "driver topup"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
