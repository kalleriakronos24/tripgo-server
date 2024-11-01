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
// @Summary      Method to give rating to the driver car
// @Description  A POST Request to give a driver's car a rating based on completed booking
// @Tags         Booking Transfer - Rating
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param 		 data body dto.InsertBookingTransferRating true "insert booking transfer rating"
// @Router       /booking/transfer/rating [post]
func POSTCreateBookingTransferRating(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.InsertBookingTransferRating{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var cred master.Customer
	if cred, err = services.Handler.RetrieveEntityCustomerByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "passenger"))
		return
	}

	p := &dto.InsertBookingTransferRating{
		CustomerID:                cred.ID,
		BookingTransferAssignedID: pValidator.BookingTransferAssignedID,
		CarManagementId:           pValidator.CarManagementId,
		Rating:                    pValidator.Rating,
	}

	if err = services.Handler.InsertBookingTransferRating(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "booking transfer rating"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
