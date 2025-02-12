package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/services"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

// AuthLogin godoc
// @Summary      Method to create company
// @Description  A POST Request to create a company
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data formData dto.InsertFormCompany true "insert company"
// @Param companyCertificate formData file true "company certificate upload"
// @Router       /company/create [post]
func POSTCreateCompany(c *gin.Context) {
	var err error

	pValidator := &dto.InsertFormCompany{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fCompanyCertificate, _ := c.FormFile("companyCertificate")
	if fCompanyCertificate == nil {
	}

	p := &dto.InsertCompany{
		Name:               pValidator.Name,
		PhoneNumber:        pValidator.PhoneNumber,
		Email:              pValidator.Email,
		CompanyName:        pValidator.CompanyName,
		CompanyAddress:     pValidator.CompanyAddress,
		CompanyCountry:     pValidator.CompanyCountry,
		CompanyNumber:      pValidator.CompanyNumber,
		CompanyCertificate: fCompanyCertificate,
		Lat:                pValidator.Lat,
		Lng:                pValidator.Lng,
	}

	pAgentPIC := &dto.DriverSignUpValidator{
		Name:  pValidator.Name,
		Email: pValidator.Email,
		Phone: pValidator.PhoneNumber,
	}

	if err = services.Handler.InsertCompany(c, p, pAgentPIC, pValidator.Drivers, pValidator.Transports); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "conpany"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to Approve company registration
// @Description  A POST Request to accept company submission request
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param        id   path      string  true  "Company ID"
// @Router       /company/create/approval/{id} [get]
func POSTApproveCompanyRegistration(c *gin.Context) {
	var err error

	// userLoggedInId := c.GetString("user_id")
	// userId, err := uuid.Parse(userLoggedInId)
	companyIdParam, _ := c.Params.Get("id")
	companyId, err := uuid.Parse(companyIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.ApproveCompany(companyId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "approve company"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to get all customer registered
// @Description  A GET Request to fetch a all records customer registered
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router       /company/customer/all [get]
func GETAllCustomerRegistered(c *gin.Context) {

	if customers, err := services.Handler.RetrieveAllRegisteredCustomerPaginated(c); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "registered customer"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &customers, Message: "success"})
	}
}

// AuthLogin godoc
// @Summary      Method to get all drivers by company id
// @Description  A GET Request to fetch a all records
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param        id   path      string  true  "Company ID"
// @Router       /company/driver/all/{id} [get]
func GETAllDriversByCompanyId(c *gin.Context) {

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	// companyIdParam, _ := c.Params.Get("id")
	// companyId, err := uuid.Parse(companyIdParam)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
	// 	return
	// }

	var driver master.Credentials
	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Driver data not found."))
		return
	}

	var companyDriver master.Driver
	if companyDriver, err = services.Handler.RetrieveDriverLinkedCompany(driver.CredentialDriver.ID); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Driver data not found."))
		return
	}

	if drivers, err := services.Handler.RetrieveAllAvailableDrivers(companyDriver.Company.ID); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "available drivers"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &drivers, Message: "success"})
	}
}

// AuthSignup godoc
// @Summary      A Driver Registration through company manager
// @Description  A Driver Registration through company manager
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param 	 	 data body dto.DriverSignup true "driver registration"
// @Router       /company/driver/new [post]
func POSTRegisteNewrDriverPartner(c *gin.Context) {
	var err error

	p := &dto.DriverSignup{}

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.Credentials }{&masterModels.Credentials{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterNewDriverInternalAgent(userId, p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "driver"))
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success"})
}
