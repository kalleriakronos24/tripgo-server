package v1

import (
	"errors"
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllProducts(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if product, err := services.Handler.RetrieveAllProduct(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: product})
	}
}

func GETAllProduct(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if product, err := services.Handler.RetrieveAllProductPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: product})
	}
}

func GETProduct(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	productId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var product models.Product
	if product, err = services.Handler.RetrieveProduct(productId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: product})
}

func POSTProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertProductValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.InsertProduct
	if pValidator.CompanyID == "" {
		var companyId uuid.UUID
		if company, err := services.Handler.RetrieveCompanyByUserID(userId); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
			return
		} else {
			companyId = company.ID
		}

		p = &dto.InsertProduct{
			Name:      pValidator.Name,
			UnitPrice: pValidator.UnitPrice,
			Packaging: pValidator.Packaging,
			Stock:     pValidator.Stock,
			Note:      pValidator.Note,
			CompanyID: companyId,
			CreatedBy: pValidator.CreatedBy,
		}
	} else {
		companyId, _ := uuid.Parse(userLoggedInId)
		p = &dto.InsertProduct{
			Name:      pValidator.Name,
			UnitPrice: pValidator.UnitPrice,
			Packaging: pValidator.Packaging,
			Stock:     pValidator.Stock,
			Note:      pValidator.Note,
			CompanyID: companyId,
			CreatedBy: pValidator.CreatedBy,
		}
	}

	if err := services.Handler.CheckExistingProduct("", struct{ *models.Product }{&models.Product{
		Name: p.Name,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with product name %s", p.Name)))
		return
	}

	if err = services.Handler.InsertProduct(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateProductValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.UpdateProduct
	if pValidator.CompanyID == "" {
		var companyId uuid.UUID
		if company, err := services.Handler.RetrieveCompanyByUserID(userId); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "company"))
			return
		} else {
			companyId = company.ID
		}

		p = &dto.UpdateProduct{
			ID:        pValidator.ID,
			Name:      pValidator.Name,
			UnitPrice: pValidator.UnitPrice,
			Packaging: pValidator.Packaging,
			Stock:     pValidator.Stock,
			Note:      pValidator.Note,
			CompanyID: companyId,
			UpdatedBy: pValidator.UpdatedBy,
		}
	} else {
		companyId, _ := uuid.Parse(pValidator.CompanyID)
		p = &dto.UpdateProduct{
			ID:        pValidator.ID,
			Name:      pValidator.Name,
			UnitPrice: pValidator.UnitPrice,
			Packaging: pValidator.Packaging,
			Stock:     pValidator.Stock,
			Note:      pValidator.Note,
			CompanyID: companyId,
			UpdatedBy: pValidator.UpdatedBy,
		}
	}

	id, _ := c.Params.Get("id")
	productId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if product, err := services.Handler.RetrieveProduct(productId); err == nil {
		if product.Name != p.Name {
			if err := services.Handler.CheckExistingProduct(id, struct{ *models.Product }{&models.Product{
				Name: p.Name,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with product name %s", p.Name)))
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product"))
		return
	}

	if err = services.Handler.UpdateProduct(productId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func DELETEProduct(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	productId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingProduct(productId.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("product id %s is not found", productId)))
		return
	}

	if err = services.Handler.DeleteProduct(productId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("delete-failed", err, "product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
