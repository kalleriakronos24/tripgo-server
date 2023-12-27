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

func GETAllCompany(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if company, err := services.Handler.RetrieveAllCompanyPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "client"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: &company})
	}

}

func GETCompany(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	companyId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var company masterModels.Company
	if company, err = services.Handler.RetrieveCompany(companyId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: company})
}
func POSTCompany(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	p := &dto.InsertCompany{CreatedBy: userId}
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

	if err := services.Handler.CheckExistingCompany("", struct{ *masterModels.Company }{&masterModels.Company{
		Email: p.Email,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, ""))
		return
	}

	if err = services.Handler.InsertCompany(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "company"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTCompany(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	companyId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	p := &dto.UpdateCompany{ID: companyId, UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.EmailFormatValidation(p.Email); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
		return
	}

	if company, err := services.Handler.RetrieveCompany(companyId); err == nil {

		if company.Email != p.Email {
			if err := services.Handler.CheckExistingCompany(id, struct{ *masterModels.Company }{&masterModels.Company{
				Email: p.Email,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing-email", err, p.Email))
				return
			}
		}
		// add another validation ?
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
		return
	}

	if err = services.Handler.UpdateCompany(companyId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "company"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func DELETECompany(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	companyId, err := uuid.Parse(id)

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingCompany(companyId.String(), struct{ *masterModels.Company }{&masterModels.Company{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("company id %s is not found", companyId)))
		return
	}

	if err = services.Handler.DeleteCompany(companyId, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("delete-failed", err, "company"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
