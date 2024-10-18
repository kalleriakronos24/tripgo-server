package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/services"
)

// AuthLogin godoc
// @Summary      List Booking Transfer Assigned
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer [get]
func GETAllBookingTransferAssignedByDriverID(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Passenger data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferAssignedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to accept booking transfer request
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/accept [get]
func POSTAcceptBookingTransfer(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Passenger data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferAssignedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to cancel booking transfer request
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/cancel [get]
func POSTCancelBookingTransfer(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Passenger data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferAssignedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}
