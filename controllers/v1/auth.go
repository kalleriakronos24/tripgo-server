package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/services"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

// AuthLogin godoc
// @Summary      Global Sign-In
// @Description  A Global authentication sign-in method for all microservices
// @Tags         Authentication - Global
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.UserLogin true "global login"
// @Router       /auth/signin [post]
func POSTLogin(c *gin.Context) {
	var err error
	var p dto.CredentialSignInDto
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var token string
	if token, err = services.Handler.AuthenticateUser(p); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect email or password. Please try again"))
		return
	}

	clientSideUrl := config.AppConfig.APPUrlClientSide

	if config.AppConfig.Environment == "DEVELOPMENT" {
		clientSideUrl = "localhost"
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, int(time.Now().Add(time.Hour*24).Unix()), "/", clientSideUrl, true, false)

	c.JSON(http.StatusCreated, dto.Response{Message: "success"})
	// if config.AppConfig.Environment == "DEVELOPMENT" {
	// 	cookie, err := c.Cookie("token")
	// 	if err != nil {
	// 		c.String(http.StatusNotFound, "Cookie not found")
	// 		return
	// 	}
	// 	c.String(http.StatusOK, "Cookie value: %s", cookie)
	// }
}

// AuthSignup godoc
// @Summary      Internal Sign-Up
// @Description  Internal Sign-Up
// @Tags         Authentication - Internal
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.UserSignup true "internal registration"
// @Router       /auth/i/signup [post]
func POSTRegister(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, _ := uuid.Parse(userLoggedInId)

	p := &dto.UserSignup{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.Credentials }{&masterModels.Credentials{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterUser(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}

// AuthSignup godoc
// @Summary      Customer Sign-Up
// @Description  Customer Sign-Up
// @Tags         Authentication - Customer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.CustomerSignup true "customer registration"
// @Router       /auth/c/signup [post]
func POSTRegisterCustomer(c *gin.Context) {
	var err error

	p := &dto.CustomerSignup{}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.Credentials }{&masterModels.Credentials{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterCustomer(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}

	var pLoginObject = dto.CredentialSignInDto{
		Email:    p.Email,
		Password: p.Password,
	}

	var token string
	if token, err = services.Handler.AuthenticateUser(pLoginObject); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect email or password. Please try again"))
		return
	}

	clientSideUrl := config.AppConfig.APPUrlClientSide

	if config.AppConfig.Environment == "DEVELOPMENT" {
		clientSideUrl = "localhost"
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, int(time.Now().Add(time.Hour*24).Unix()), "/", clientSideUrl, true, false)
	c.SetCookie("name", p.Name, int(time.Now().Add(time.Hour*24).Unix()), "/", clientSideUrl, true, false)

	c.JSON(http.StatusCreated, dto.Response{Message: "success"})
}

// AuthSignup godoc
// @Summary      Driver Sign-Up
// @Description  Driver Sign-Up
// @Tags         Authentication - Driver
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.DriverSignup true "driver registration"
// @Router       /auth/d/signup [post]
func POSTRegisterDriver(c *gin.Context) {
	var err error

	p := &dto.DriverSignup{}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.Credentials }{&masterModels.Credentials{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterDriver(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "driver"))
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}
