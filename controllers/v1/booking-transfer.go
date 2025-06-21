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
// @Summary      Method to create booking transfer request by customer
// @Description  A POST Request to create a new record for costumer to make a new Transfer booking to the driver
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.InsertBookingTransfer true "insert booking transfer"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking/transfer [post]
func POSTBookingTransfer(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, _ := uuid.Parse(userLoggedInId)

	pValidator := &dto.InsertBookingTransfer{}
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
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	p := &dto.InsertBookingTransfer{
		AdultSeater:       pValidator.AdultSeater,
		ChildSeater:       pValidator.ChildSeater,
		FromLatCoordinate: pValidator.FromLatCoordinate,
		FromLngCoordinate: pValidator.FromLngCoordinate,
		ToLatCoordinate:   pValidator.ToLatCoordinate,
		ToLngCoordinate:   pValidator.ToLngCoordinate,
		FromLocation:      pValidator.FromLocation,
		ToLocation:        pValidator.ToLocation,
		PassengerNotes:    pValidator.PassengerNotes,
		PickUpDate:        pValidator.PickUpDate,
		CarModelID:        pValidator.CarModelID,
		Price:             pValidator.Price,
		GrandTotal:        pValidator.GrandTotal,
		AddPickupPoint:    pValidator.AddPickupPoint,
		AddDropoffPoint:   pValidator.AddDropoffPoint,
		CustomerID:        cred.CredentialCustomer.ID,
		TotalDistance:     pValidator.TotalDistance,
		RefferalCode:      pValidator.RefferalCode,
		PaymentOption:     pValidator.PaymentOption,
		PI:                pValidator.PI,
		Currency:          pValidator.Currency,
	}

	// var bookingTransfer models.BookingTransfer
	// if bookingTransfer, err = services.Handler.RetrieveLastOrderByCustomerID(cred.CredentialCustomer.ID); err != nil {
	// 	if err = services.Handler.InsertBookingTransfer(p); err != nil {
	// 		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "booking transfer"))
	// 		return
	// 	}
	// 	c.JSON(http.StatusCreated, dto.Response{Message: "success"})
	// 	// c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
	// 	return
	// }

	// if bookingTransfer.IsCompleted == utils.NewFalse() {
	// 	c.JSON(http.StatusBadRequest, constants.GetErrorResponse("general", nil, "You can only have one active booking."))
	// 	return
	// }

	if err = services.Handler.InsertBookingTransfer(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "booking transfer"))
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to get all booking transfer made by customer
// @Description  A GET Request to fetch a all records for customer to view booking transfer that already created
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking/transfer/all [get]
func GETAllBookingTransferByCustomer(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var cred master.Customer
	if cred, err = services.Handler.RetrieveEntityCustomerByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	if bookingTransfer, err := services.Handler.RetrieveAllBookingTransferByCustomer(cred.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransfer, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to get all received khaimal booking transfers
// @Description  A GET Request to fetch a all records booking transfers
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Router       /booking/transfer/khaimal [get]
func GETAllKhaimalBookingTransfer(c *gin.Context) {
	if bookingTransfer, err := services.Handler.RetrieveAllKhaimalBookingTransfer(); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransfer, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to get all received partner booking transfers
// @Description  A GET Request to fetch a all records booking transfers
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /booking/transfer/partner [get]
func GETAllPartnerBookingTransfer(c *gin.Context) {
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "internal"))
		return
	}

	if bookingTransfer, err := services.Handler.RetrieveAllPartnerBookingTransfer(cred.CredentialInternal.RefferalCode); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransfer, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to count booking transfer by customer
// @Description  A GET Request to fetch count how many bookings are made by customer either transfer, tour or delivery
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /web/statistic/customer/booking-count [get]
func GETCountBookingTransferByCustomer(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var cred master.Customer
	if cred, err = services.Handler.RetrieveEntityCustomerByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	if bookingTransfer, err := services.Handler.RetrieveCustomerWebStatisticByUserID(cred.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransfer, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to count active booking transfer by customer
// @Description  A GET Request to fetch active count how many bookings are made by customer either transfer, tour or delivery
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /web/statistic/customer/booking-count/active [get]
func GETActiveCountBookingTransferByCustomer(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var cred master.Customer
	if cred, err = services.Handler.RetrieveEntityCustomerByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	if bookingTransfer, err := services.Handler.RetrieveCustomerWebStatisticActiveByUserID(cred.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &bookingTransfer, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to cancel booking transfer by customer
// @Description  A POST Request to cancel the selected booking transfer and notify the assigned driver
// @Tags         Booking - Transfer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "Booking Transfer Assigned ID"
// @Router       /booking/transfer/cancel/{id} [post]
func POSTCancelBookingTransferByCustomer(c *gin.Context) {
	var err error

	bookingTransferAssignedIdParam, _ := c.Params.Get("id")
	bookingTransferAssignedId, err := uuid.Parse(bookingTransferAssignedIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CustomerCancelBooking(bookingTransferAssignedId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
