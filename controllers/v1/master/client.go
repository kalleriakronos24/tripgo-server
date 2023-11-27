package v1

import (
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllClient(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var client []masterModels.Client
	if client, err = services.Handler.RetrieveAllClient(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: client})
}

func GETClient(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	clientId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var client masterModels.Client
	if client, err = services.Handler.RetrieveClient(clientId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: client})
}

func POSTClient(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.InsertClient{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, ""))
		return
	}

	if err := services.Handler.CheckExistingClient("", struct{ *masterModels.Client }{&masterModels.Client{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.InsertClient(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "client"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTClient(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.UpdateClient{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	clientId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if client, err := services.Handler.RetrieveClient(clientId); err == nil {

		if client.Email != p.Email {
			if err := services.Handler.CheckExistingClient(id, struct{ *masterModels.Client }{&masterModels.Client{
				Email: p.Email,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, p.Email))
				return
			}
		}

		if client.Name != p.Name {
			if err := services.Handler.CheckExistingClient(id, struct{ *masterModels.Client }{&masterModels.Client{
				Name: p.Name,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-username", err, p.Name))
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	}

	if err = services.Handler.UpdateClient(clientId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "client"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
