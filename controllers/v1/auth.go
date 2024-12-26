package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	middleware "github.com/kalleriakronos24/khaimal-group/controllers/middlewares"
	"github.com/kalleriakronos24/khaimal-group/dto"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/services"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

// AuthLogin godoc
// @Summary      Customer Sign-In
// @Description  A Customer authentication sign-in method
// @Tags         Authentication - Customer
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.UserLogin true "customer login"
// @Router       /auth/c/signin [post]
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

	var customer masterModels.Credentials
	claims, err := middleware.ParseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: "Unauthorized. Please Login Again"})
		return
	}

	if customer, err = services.Handler.RetrieveEntityCredentialsByUserID(claims.ID); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Passenger data not found."))
		return
	}

	responseData := struct {
		Token    string `json:"token,omitempty"`
		FullName string `json:"fullName,omitempty"`
	}{
		Token:    token,
		FullName: customer.CredentialCustomer.Name,
	}

	// TODO
	// CHANGE TO COOKIE LATER
	// clientSideUrl := config.AppConfig.APPUrlClientSide

	// if config.AppConfig.Environment == "DEVELOPMENT" {
	// 	clientSideUrl = "https://khaimal-webprofile.vercel.app"
	// }

	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("token", token, int(time.Now().Add(time.Hour*24).Unix()), "", clientSideUrl, true, false)

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: responseData})
}

// AuthLogin godoc
// @Summary      Driver Sign-In
// @Description  A Driver authentication sign-in method for mobile apps
// @Tags         Authentication - Driver
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.UserLogin true "driver login"
// @Router       /auth/d/signin [post]
func POSTLoginDriver(c *gin.Context) {
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

	var driver masterModels.Credentials
	claims, err := middleware.ParseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: "Unauthorized. Please Login Again"})
		return
	}

	if driver, err = services.Handler.RetrieveEntityCredentialsByUserID(claims.ID); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "Passenger data not found."))
		return
	}

	responseData := struct {
		Token    string `json:"token,omitempty"`
		FullName string `json:"fullName,omitempty"`
		Type     string `json:"type,omitempty"`
	}{
		Token:    token,
		FullName: driver.CredentialDriver.Name,
		Type:     driver.CredentialDriver.DriverType,
	}

	// TODO
	// CHANGE TO COOKIE LATER
	// clientSideUrl := config.AppConfig.APPUrlClientSide

	// if config.AppConfig.Environment == "DEVELOPMENT" {
	// 	clientSideUrl = "https://khaimal-webprofile.vercel.app"
	// }

	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("token", token, int(time.Now().Add(time.Hour*24).Unix()), "", clientSideUrl, true, false)

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: responseData})
}

// AuthLogin godoc
// @Summary      Internal Sign-In
// @Description  A Internal authentication sign-in method
// @Tags         Authentication - Internal
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 		 data body dto.CredentialSignInDto true "internal login"
// @Router       /auth/i/signin [post]
func POSTLoginInternal(c *gin.Context) {
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

	var userInternal masterModels.Credentials
	claims, err := middleware.ParseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Message: "Unauthorized. Please Login Again"})
		return
	}

	if userInternal, err = services.Handler.RetrieveEntityCredentialsByUserID(claims.ID); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("data-not-found", err, "User internal data not found."))
		return
	}

	responseData := struct {
		Token    string `json:"token,omitempty"`
		FullName string `json:"fullName,omitempty"`
	}{
		Token:    token,
		FullName: userInternal.CredentialInternal.Name,
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: responseData})
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

	p := &dto.UserSignup{}

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

	responseData := struct {
		Token    string `json:"token,omitempty"`
		FullName string `json:"fullName,omitempty"`
	}{
		Token:    token,
		FullName: p.Name,
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: responseData})
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

	var pLoginObject = dto.CredentialSignInDto{
		Email:    p.Email,
		Password: p.Password,
	}

	var token string
	if token, err = services.Handler.AuthenticateUser(pLoginObject); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect email or password. Please try again"))
		return
	}

	responseData := struct {
		Token    string `json:"token,omitempty"`
		FullName string `json:"fullName,omitempty"`
	}{
		Token:    token,
		FullName: p.Name,
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: responseData})
}

// AuthSignup godoc
// @Summary      Universal Reset Password
// @Description  Universal Reset Password
// @Tags         Authentication - Reset Password
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.DriverSignup true "universal reset password"
// @Router       /auth/reset-password [post]
func POSTResetPassword(c *gin.Context) {
	var err error

	p := &dto.UniversalResetPassword{}

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
	}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "email not found"))
		return
	}

	if err = services.Handler.ResetPassword(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "success", Data: nil})
}
