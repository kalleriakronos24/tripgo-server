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

// AuthLogin godoc
// @Summary      Method to get all car management made by driver
// @Description  A GET Request to fetch a all records for driver to view list of added car management that already created
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /car-management/get [get]
func GETAllCarManagementByDriverID(c *gin.Context) {
	var err error

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

	if carManagements, err := services.Handler.RetrieveCarManagementrByUserID(cred.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "car management"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &carManagements, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to update car management made by driver
// @Description  A UPDATE Request to update record car management
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data formData dto.UpdateFormCarManagement true "update car management"
// @Param frontCarPhoto formData file true "front car photo upload"
// @Router       /car-management/update/{id} [put]
func PUTUpdateCarManagementByID(c *gin.Context) {
	var err error

	carManagementIdParam, _ := c.Params.Get("id")
	carManagementId, err := uuid.Parse(carManagementIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.UpdateFormCarManagement{}
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

	carModelid, _ := uuid.Parse(pValidator.CarModelID)
	p := &dto.UpdateCarManagement{
		Name:          pValidator.Name,
		PlateNumber:   pValidator.PlateNumber,
		LicensePhoto:  pValidator.LicensePhoto,
		CarModelID:    carModelid,
		FrontCarPhoto: fFrontCarPhotoDocument,
	}

	if err := services.Handler.UpdateCarManagement(c, p, carManagementId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "car management"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: nil, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to delete car management made by driver
// @Description  A DELETE Request to delete record car management
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /car-management/delete/{id} [delete]
func DELCarManagementByID(c *gin.Context) {
	var err error

	carManagementIdParam, _ := c.Params.Get("id")
	carManagementId, err := uuid.Parse(carManagementIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.DeleteCarManagement(c, carManagementId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "car management"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: nil, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set inactive car management made by driver
// @Description  A UPDATE Request to set record car management to inactive
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /car-management/set/inactive/{id} [put]
func UPDCarManagementToInactiveByID(c *gin.Context) {
	var err error

	carManagementIdParam, _ := c.Params.Get("id")
	carManagementId, err := uuid.Parse(carManagementIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.SetCarManagementInactive(c, carManagementId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "car management"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: nil, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to set active car management made by driver
// @Description  A UPDATE Request to set record car management to active
// @Tags         Car Management
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /car-management/set/active/{id} [put]
func UPDCarManagementToActiveByID(c *gin.Context) {
	var err error

	carManagementIdParam, _ := c.Params.Get("id")
	carManagementId, err := uuid.Parse(carManagementIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.SetCarManagementActive(c, carManagementId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "car management"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: nil, Message: "success"})
}
