package v1

import (
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

func GETAllProductHistory(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if ProductHistory, err := services.Handler.RetrieveAllProductHistoryPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "ProductHistory"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: ProductHistory})
	}

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
	pValidator := &dto.InsertProductHistoryValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	productId, _ := uuid.Parse(pValidator.ProductID)
	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)

	p := &dto.InsertProductHistory{
		Status:              pValidator.Status,
		Quantity:            pValidator.Quantity,
		ProductID:           productId,
		OperatingActivityID: operatingActivityId,
		CreatedBy:           pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("product id %s is not found", p.ProductID)))
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
	pValidator := &dto.UpdateProductHistoryValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	productId, _ := uuid.Parse(pValidator.ProductID)
	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)

	p := &dto.UpdateProductHistory{
		ID:                  pValidator.ID,
		Status:              pValidator.Status,
		Quantity:            pValidator.Quantity,
		ProductID:           productId,
		OperatingActivityID: operatingActivityId,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	ProductHistoryId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("product id %s is not found", p.ProductID)))
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
