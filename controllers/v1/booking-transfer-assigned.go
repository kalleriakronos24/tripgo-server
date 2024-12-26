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
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Driver data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferAssignedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "assigned booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      List Booking Transfer Assigned
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/accepted [get]
func GETAllBookingTransferAccepteddByDriverID(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Driver data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferAcceptedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "accepted booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      List Booking Transfer Assigned
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/cancelled [get]
func GETAllBookingTransferCancelledByDriverID(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Driver data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferCancelledByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "cancelled booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      List Booking Transfer Assigned
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/ongoing [get]
func GETAllBookingTransferOngoingByDriverID(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Driver data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferOngoingByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "ongoing booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      List Booking Transfer Assigned
// @Description  A GET Request to fetch a all records for driver to view booking transfer that assigned and whose are not accepted by the driver
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking-assigned/transfer/completed [get]
func GETAllBookingTransferCompletedByDriverID(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Driver data not found."))
		return
	}

	if bookingTransferAssigned, err := services.Handler.RetrieveBookingTransferCompletedByDriverID(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "completed booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransferAssigned, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to accept booking transfer
// @Description  A POST Request for Driver to Accept specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking-assigned/transfer/accept/{id} [get]
func POSTAcceptBookingTransfer(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	// userId, err := uuid.Parse(userLoggedInId)
	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	// var driver master.Credentials
	// if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
	// 	c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Driver data not found."))
	// 	return
	// }

	if err := services.Handler.AcceptBookingTransfer(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "accept booking transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to cancel booking transfer request
// @Description  A POST Request for Driver to Cancel specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking-assigned/transfer/cancel [get]
func POSTCancelBookingTransfer(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	// userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}
	// todo
	// var driver master.Credentials
	// if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
	// 	c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Passenger data not found."))
	// 	return
	// }

	if err := services.Handler.CancelBookingTransfer(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "cancel booking transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set status to ongoing and notify passenger that driver on it's way to pickup
// @Description  A POST Request for Driver to set status to be "OnGoing" to specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking-assigned/transfer/ongoing/{id} [get]
func POSTOngoingBookingTransfer(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	// userId, err := uuid.Parse(userLoggedInId)

	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	// var driver master.Credentials
	// if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
	// 	c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Driver data not found."))
	// 	return
	// }

	if err := services.Handler.OngoingBookingTransfer(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "ongoing booking transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set status to complete and notify passenger
// @Description  A POST Request for Driver to set Complete specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking-assigned/transfer/ongoing [get]
func POSTCompleteBookingTransfer(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	// userId, err := uuid.Parse(userLoggedInId)

	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CompleteBookingTransfer(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "complete booking transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set status to Complete PickUp and notify passenger
// @Description  A POST Request for Driver to set Complete Pickup specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking-assigned/transfer/ongoing [get]
func POSTPickupBookingTransfer(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	// userId, err := uuid.Parse(userLoggedInId)

	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CompletePickupBooking(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "complete pickup transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set status to Complete PickUp and notify passenger
// @Description  A POST Request for Driver to set Complete Pickup specific booking transfer by it's ID
// @Tags         Booking Assgined - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param        id   path      string  true  "BookingTransferAssigned ID"
// @Param        driverId   path      string  true  "DriverID ID"
// @Param        plateNumber   path      string  true  "Car Plate Number"
// @Router       /booking-assigned/transfer/switch/{id}/{driverId}/{plateNumber} [get]
func POSTSwitchDriverByAgent(c *gin.Context) {
	var err error

	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	driverIdParamIdParam, _ := c.Params.Get("driverId")
	driverId, err := uuid.Parse(driverIdParamIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	plateNumberParam, _ := c.Params.Get("plateNumber")

	if err := services.Handler.SwitchDriver(bookingTransferAssignedId, driverId, plateNumberParam); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "complete pickup transfer order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
