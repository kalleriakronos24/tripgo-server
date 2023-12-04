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

func GETAllOperatingActivityProduct(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if operatingActivityProduct, err := services.Handler.RetrieveAllOperatingActivityProductPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity product"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: operatingActivityProduct})
	}

}

func GETOperatingActivityProduct(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	operatingActivityProductId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var operatingActivityProduct models.OperatingActivityProduct
	if operatingActivityProduct, err = services.Handler.RetrieveOperatingActivityProduct(operatingActivityProductId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity product"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: operatingActivityProduct})
}

func POSTOperatingActivityProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertOperatingActivityProductValidator{CreatedBy: userId}

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
	p := &dto.InsertOperatingActivityProduct{
		Quantity:            pValidator.Quantity,
		VATRate:             pValidator.VATRate,
		SubTotal:            pValidator.SubTotal,
		GrandTotal:          pValidator.GrandTotal,
		OperatingActivityID: operatingActivityId,
		ProductID:           productId,
		CreatedBy:           pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivityProduct("", struct {
		*models.OperatingActivityProduct
	}{&models.OperatingActivityProduct{
		OperatingActivityID: p.OperatingActivityID,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with operating activity id %s", p.OperatingActivityID)))
		return
	}

	if err = services.Handler.InsertOperatingActivityProduct(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "operating activity id"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTOperatingActivityProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateOperatingActivityProductValidator{UpdatedBy: userId}

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
	p := &dto.UpdateOperatingActivityProduct{
		Quantity:            pValidator.Quantity,
		VATRate:             pValidator.VATRate,
		SubTotal:            pValidator.SubTotal,
		GrandTotal:          pValidator.GrandTotal,
		OperatingActivityID: operatingActivityId,
		ProductID:           productId,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	operatingActivityProductId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if operatingActivityProduct, err := services.Handler.RetrieveOperatingActivityProduct(operatingActivityProductId); err == nil {

		if operatingActivityProduct.OperatingActivityID != p.OperatingActivityID {
			if err := services.Handler.CheckExistingOperatingActivityProduct(id, struct {
				*models.OperatingActivityProduct
			}{&models.OperatingActivityProduct{
				OperatingActivityID: p.OperatingActivityID,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with operating activity's tax number %s", operatingActivityProduct.OperatingActivity.TaxInvoiceNumber)))
				return
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "operating activity product"))
		return
	}

	if err = services.Handler.UpdateOperatingActivityProduct(operatingActivityProductId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "operating activity product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
