package v1

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/utils"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/services"
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
	var p dto.UserLogin
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

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, int(time.Now().Add(time.Hour*24).Unix()), "", "", true, false)
	if config.AppConfig.Environment == "DEVELOPMENT" {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusNotFound, "Cookie not found")
			return
		}
		c.String(http.StatusOK, "Cookie value: %s", cookie)
	}
}

// AuthSignup godoc
// @Summary      User Sign-Up
// @Description  User Sign-Up to store them into global database
// @Tags         Authentication - Global
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.UserSignup true "user registration"
// @Router       /auth/signup [post]
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
