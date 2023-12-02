package v1

import (
	"errors"
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GETAllPurchaseOrderProduct(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PurchaseOrderProduct []models.PurchaseOrderProduct
	if PurchaseOrderProduct, err = services.Handler.RetrieveAllPurchaseOrderProduct(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order product"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: PurchaseOrderProduct})
}

func GETPurchaseOrderProduct(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	PurchaseOrderProductId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PurchaseOrderProduct models.PurchaseOrderProduct
	if PurchaseOrderProduct, err = services.Handler.RetrievePurchaseOrderProduct(PurchaseOrderProductId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: PurchaseOrderProduct})
}

func POSTPurchaseOrderProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.InsertPurchaseOrderProduct{CreatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("product id %s is not found", p.ProductID)))
		return
	}

	if err := services.Handler.CheckExistingClient(p.ClientID.String(), struct{ *masterModels.Client }{&masterModels.Client{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("client id %s is not found", p.ClientID)))
		return
	}

	if err := services.Handler.CheckExistingPurchaseOrder(p.PurchaseOrderID.String(), struct{ *models.PurchaseOrder }{&models.PurchaseOrder{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("purchase order id %s is not found", p.ClientID)))
		return
	}

	if err := services.Handler.CheckExistingPurchaseOrderProduct("", struct {
		*models.PurchaseOrderProduct
	}{&models.PurchaseOrderProduct{
		PurchaseOrderID: p.PurchaseOrderID,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with purchase order id %s", p.PurchaseOrderID.String())))
		return
	}

	if err = services.Handler.InsertPurchaseOrderProduct(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "purchase order product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTPurchaseOrderProduct(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.UpdatePurchaseOrderProduct{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}
	id, _ := c.Params.Get("id")
	PurchaseOrderProductId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingProduct(p.ProductID.String(), struct{ *models.Product }{&models.Product{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("product id %s is not found", p.ProductID)))
		return
	}

	if err := services.Handler.CheckExistingClient(p.ClientID.String(), struct{ *masterModels.Client }{&masterModels.Client{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("client id %s is not found", p.ClientID)))
		return
	}

	if err := services.Handler.CheckExistingPurchaseOrder(p.PurchaseOrderID.String(), struct{ *models.PurchaseOrder }{&models.PurchaseOrder{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("purchase order id %s is not found", p.ClientID)))
		return
	}

	if PurchaseOrderProduct, err := services.Handler.RetrievePurchaseOrderProduct(PurchaseOrderProductId); err == nil {

		if PurchaseOrderProduct.PurchaseOrderID != p.PurchaseOrderID {
			if err := services.Handler.CheckExistingPurchaseOrderProduct(id, struct {
				*models.PurchaseOrderProduct
			}{&models.PurchaseOrderProduct{
				PurchaseOrderID: p.PurchaseOrderID,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with purchase order number %s", PurchaseOrderProduct.PurchaseOrder.Number)))
				return
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order product"))
		return
	}

	if err = services.Handler.UpdatePurchaseOrderProduct(PurchaseOrderProductId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "purchase order product"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
