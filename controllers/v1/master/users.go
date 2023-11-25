package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/services"
)

func GETUser(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	log.Println(id)
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: "Parameter ID is empty"})
		return
	}

	var user masterModels.User
	if user, err = services.Handler.RetrieveUser(userId); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Error: "user not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: dto.RetrieveUserInfo{
		Name:     user.Name,
		Address:  user.Address,
		Username: user.Username,
		Email:    user.Email,
	}})
}

func PUTUser(c *gin.Context) {
	var err error
	var user dto.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: "Parameter ID is empty"})
		return
	}
	if err = services.Handler.UpdateUser(userId, user); err != nil {
		c.JSON(http.StatusNotModified, dto.Response{Error: "failed updating user"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: user})
}
