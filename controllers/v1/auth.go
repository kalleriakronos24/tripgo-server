package v1

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/constants"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func POSTLogin(c *gin.Context) {
	var err error
	var p dto.UserLogin
	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	log.Println("is this working")

	var token string
	if token, err = services.Handler.AuthenticateUser(p); err != nil {
		c.JSON(http.StatusNotFound, constants.GetErrorResponse("logical", err, "Incorrect username or password. Please try again"))
		return
	}

	log.Println("???")

	c.JSON(http.StatusOK, dto.Response{Data: token})
}

func POSTRegister(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	log.Println(userId)
	p := &dto.UserSignup{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fmt.Printf("value >>>>> %v", p)

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

func POSTRegisterSuperAdmin(c *gin.Context) {
	var err error

	var p dto.UserSignupSuperAdmin
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

	if err = services.Handler.RegisterUserSuperAdmin(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "user"))
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "success"})
}
