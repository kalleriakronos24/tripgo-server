package v1

import (
	"errors"
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
// @Summary      Method to create or assign a car to the driver
// @Description  A POST Request to create a car and assigned to the driver whose creating it
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data formData dto.InsertFormCarManagement true "insert car management"
// @Param frontCarPhoto formData file true "front car photo upload"
// @Router       /car-management/create [post]
func POSTCreateCarManagement(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.InsertFormCarManagement{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fFrontCarPhotoDocument, _ := c.FormFile("frontCarPhoto")
	if fFrontCarPhotoDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field frontCarPhoto is required"), "cannot submit if front car photo is empty"))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "driver"))
		return
	}

	carModelid, _ := uuid.Parse(pValidator.CarModelID)
	p := &dto.InsertCarManagement{
		Name:              pValidator.Name,
		PlateNumber:       pValidator.PlateNumber,
		LicensePhoto:      pValidator.LicensePhoto,
		CarManagementType: "internal",
		CarModelID:        carModelid,
		DriverID:          cred.CredentialDriver.ID,
		FrontCarPhoto:     fFrontCarPhotoDocument,
	}

	if err = services.Handler.InsertCarManagement(c, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "car management"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
