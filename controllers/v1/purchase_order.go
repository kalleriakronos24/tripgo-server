package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/services"
	"net/http"
)

func GETAllPurchaseOrder(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PurchaseOrder []models.PurchaseOrder
	if PurchaseOrder, err = services.Handler.RetrieveAllPurchaseOrder(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: PurchaseOrder})
}

func GETPurchaseOrder(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	PurchaseOrderId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var PurchaseOrder models.PurchaseOrder
	if PurchaseOrder, err = services.Handler.RetrievePurchaseOrder(PurchaseOrderId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: PurchaseOrder})
}

func POSTPurchaseOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	fPurchaseOrderDocument, err := c.FormFile("document")

	p := &dto.InsertFormDataPurchaseOrder{CreatedBy: userId}

	//var data *multipart.Form
	if err = c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityID, err := uuid.Parse(p.OperatingActivityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "invalid operating activity id"))
		return
	}

	repopulateFormDataPayload := &dto.InsertPurchaseOrder{
		Number:              p.Number,
		Type:                p.Type,
		Recipient:           p.Recipient,
		RecipientEmail:      p.RecipientEmail,
		Date:                p.Date,
		Document:            fPurchaseOrderDocument,
		OperatingActivityID: operatingActivityID,
		CreatedBy:           p.CreatedBy,
	}

	fmt.Printf("%v", repopulateFormDataPayload)

	if err := services.Handler.CheckExistingPurchaseOrder("", struct {
		*models.PurchaseOrder
	}{&models.PurchaseOrder{
		OperatingActivityID: operatingActivityID,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, ""))
		return
	}

	if err = services.Handler.InsertPurchaseOrder(c, repopulateFormDataPayload); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "purchase order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTPurchaseOrder(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	p := &dto.UpdateFormDataPurchaseOrder{UpdatedBy: userId}

	if err = c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	id, _ := c.Params.Get("id")
	PurchaseOrderId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	operatingActivityID, err := uuid.Parse(p.OperatingActivityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "invalid operating activity id"))
		return
	}

	repopulateFormDataPayload := &dto.UpdatePurchaseOrder{
		Number:              p.Number,
		Type:                p.Type,
		Recipient:           p.Recipient,
		RecipientEmail:      p.RecipientEmail,
		Date:                p.Date,
		OperatingActivityID: operatingActivityID,
	}

	if PurchaseOrder, err := services.Handler.RetrievePurchaseOrder(PurchaseOrderId); err == nil {

		if PurchaseOrder.OperatingActivityID != repopulateFormDataPayload.OperatingActivityID {
			if err := services.Handler.CheckExistingPurchaseOrder(id, struct {
				*models.PurchaseOrder
			}{&models.PurchaseOrder{
				OperatingActivityID: repopulateFormDataPayload.OperatingActivityID,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-existing", err, PurchaseOrder.OperatingActivity.DeliveryReceiptNumber))
				return
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "purchase order"))
		return
	}

	if err = services.Handler.UpdatePurchaseOrder(c, PurchaseOrderId, repopulateFormDataPayload); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "purchase order"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
