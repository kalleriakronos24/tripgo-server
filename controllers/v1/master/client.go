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

func GETAllClients(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if client, err := services.Handler.RetrieveAllClient(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &client})
	}
}

func GETAllClient(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if client, err := services.Handler.RetrieveAllClientPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &client})
	}
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
	pValidator := &dto.InsertClientValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.InsertClient
	if pValidator.CompanyID == "" {
		var companyId uuid.UUID
		if company, err := services.Handler.RetrieveCompanyByUserID(userId); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
			return
		} else {
			companyId = company.ID
		}
		p = &dto.InsertClient{
			Name:        pValidator.Name,
			PhoneNumber: pValidator.PhoneNumber,
			Email:       pValidator.Email,
			Address:     pValidator.Address,
			CreatedBy:   pValidator.CreatedBy,
			CompanyID:   companyId,
		}
	} else {
		companyId, _ := uuid.Parse(pValidator.CompanyID)
		p = &dto.InsertClient{
			Name:        pValidator.Name,
			PhoneNumber: pValidator.PhoneNumber,
			Email:       pValidator.Email,
			Address:     pValidator.Address,
			CreatedBy:   pValidator.CreatedBy,
			CompanyID:   companyId,
		}
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

	id, _ := c.Params.Get("id")
	clientId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	pValidator := &dto.UpdateClientValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.UpdateClient
	if pValidator.CompanyID == "" {
		var companyId uuid.UUID
		if company, err := services.Handler.RetrieveCompanyByUserID(userId); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
			return
		} else {
			companyId = company.ID
		}

		p = &dto.UpdateClient{
			Name:        pValidator.Name,
			PhoneNumber: pValidator.PhoneNumber,
			Email:       pValidator.Email,
			Address:     pValidator.Address,
			UpdatedBy:   pValidator.UpdatedBy,
			CompanyID:   companyId,
		}
	} else {
		companyId, _ := uuid.Parse(pValidator.CompanyID)
		p = &dto.UpdateClient{
			Name:        pValidator.Name,
			PhoneNumber: pValidator.PhoneNumber,
			Email:       pValidator.Email,
			Address:     pValidator.Address,
			UpdatedBy:   pValidator.UpdatedBy,
			CompanyID:   companyId,
		}
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
