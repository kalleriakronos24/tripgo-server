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

func GETAllProductHistory(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var ProductHistory []models.ProductHistory
	if ProductHistory, err = services.Handler.RetrieveAllProductHistory(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "ProductHistory"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: ProductHistory})
}

func GETProductHistory(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	ProductHistoryId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var ProductHistory models.ProductHistory
	if ProductHistory, err = services.Handler.RetrieveProductHistory(ProductHistoryId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product history"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: ProductHistory})
}

func POSTProductHistory(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.InsertProductHistory{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err = services.Handler.InsertProductHistory(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "product history"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTProductHistory(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.UpdateProductHistory{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	ProductHistoryId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if _, err := services.Handler.RetrieveProductHistory(ProductHistoryId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "product history"))
		return
	}

	if err = services.Handler.UpdateProductHistory(ProductHistoryId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "product history"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
