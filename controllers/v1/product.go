package v1

import (
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllProduct(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var product []models.Product
	if product, err = services.Handler.RetrieveAllProduct(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: product})
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
	p := &dto.InsertProduct{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingProduct("", struct{ *models.Product }{&models.Product{
		Name: p.Name,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, ""))
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
	p := &dto.UpdateProduct{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
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
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, p.Name))
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
