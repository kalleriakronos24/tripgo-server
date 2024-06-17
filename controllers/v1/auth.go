package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/booklap-be/constants"
	masterModels "github.com/kalleriakronos24/booklap-be/models/master"
	"github.com/kalleriakronos24/booklap-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/booklap-be/dto"
	"github.com/kalleriakronos24/booklap-be/services"
)

// AuthLogin godoc
// @Summary      Global Sign-In
// @Description  A Global authentication sign-in method for all microservices
// @Tags         Authentication
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

	c.JSON(http.StatusOK, dto.Response{Data: token})
}

// AuthSignup godoc
// @Summary      User Sign-Up
// @Description  User Sign-Up to store them into global database
// @Tags         Authentication
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

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.User }{&masterModels.User{
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

// AuthSignUp godoc
// @Summary      Internal Sign-Up
// @Description  Internal Sign-Up
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.InternalSignUp true "internal registration"
// @Router       /auth/i/signup [post]
func POSTRegisterInternal(c *gin.Context) {
	var err error

	var p dto.InternalSignUp
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.User }{&masterModels.User{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterInternal(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}

// TENANT SIGN-UP

// AuthSignup godoc
// @Summary      Owner Sign-Up
// @Description  Owner Sign-Up to store them into global database
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.OwnerSignup true "owner registration"
// @Router       /auth/t/signup [post]
func POSTRegisterBusinessOwner(c *gin.Context) {
	var err error

	p := &dto.OwnerSignup{}

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

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.User }{&masterModels.User{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterOwner(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}

// AuthSignup godoc
// @Summary      Admin Sign-Up
// @Description  Admin Sign-Up to store them into global database
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param 	 	 data body dto.AdminSignup true "admin registration"
// @Router       /auth/t/signup/admin [post]
func POSTRegisterBusinessAdmin(c *gin.Context) {
	var err error

	p := &dto.AdminSignup{}

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

	if err := services.Handler.CheckExistingUser("", struct{ *masterModels.User }{&masterModels.User{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.RegisterAdmin(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}

// TENTANT SIGN-IN
func POSTLoginBusinessOwner(c *gin.Context) {
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
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect username or password. Please try again"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: token})
}
func POSTLoginBusinessAdmin(c *gin.Context) {
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
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect username or password. Please try again"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: token})
}
