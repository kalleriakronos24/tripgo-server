package v1

import (
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/services"
)

func GETUserByID(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var user masterModels.User
	if user, err = services.Handler.RetrieveUser(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "user"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: user})
}

func GETUser(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	var user masterModels.User
	if user, err = services.Handler.RetrieveUser(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "user"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: user})
}

func PUTUser(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	p := dto.UserUpdate{ID: userId}
	if err = c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	if err := utils.ValidateHTTPPayload(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if user, err := services.Handler.RetrieveUser(userId); err == nil {

		if user.Email != p.Email {
			if err := services.Handler.CheckExistingUser(id, struct{ *masterModels.User }{&masterModels.User{
				Email: p.Email,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, p.Email))
				return
			}
		}

		if user.Username != p.Username {
			if err := services.Handler.CheckExistingUser(id, struct{ *masterModels.User }{&masterModels.User{
				Username: p.Username,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-username", err, p.Username))
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "user"))
		return
	}

	if err := services.Handler.UpdateUser(userId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "user"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func DELETEUser(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingUser(userId.String(), struct{ *masterModels.User }{&masterModels.User{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("user id %s is not found", userId)))
		return
	}

	if err = services.Handler.DeleteUser(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("delete-failed", err, "user"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
