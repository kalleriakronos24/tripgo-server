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
// @Summary      Method to get wallet information by driver
// @Description  A GET Request to fetch a record by driver to check their wallet information
// @Tags         Driver Balance
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /driver-balance/get [get]
func GETDWalletInformationByDriver(c *gin.Context) {

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

	if driverWallet, err := services.Handler.RetrieveDriverBalanceDetailByDriver(cred.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "wallet"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &driverWallet, Message: "success"})
	}
}
