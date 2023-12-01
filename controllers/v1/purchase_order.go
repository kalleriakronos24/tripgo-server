package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/services"
	"gitlab.com/odma1/odma-be/utils"
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

	var PurchaseOrder models.CustomResponsePurchaseOrder
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

	fPurchaseOrderDocument, _ := c.FormFile("document")
	if fPurchaseOrderDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field document is required"), "cannot submit if document is empty"))
		return
	}

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

	if err := utils.EmailFormatValidation(p.RecipientEmail); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email format"))
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

	// check existing operating id
	if err := services.Handler.CheckExistingPurchaseOrder("", struct {
		*models.PurchaseOrder
	}{&models.PurchaseOrder{
		OperatingActivityID: operatingActivityID,
	}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	// check duplication operating id
	if err := services.Handler.CheckExistingPurchaseOrder("", struct {
		*models.PurchaseOrder
	}{&models.PurchaseOrder{
		OperatingActivityID: operatingActivityID,
	}}); err == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with operating activity id %s", p.OperatingActivityID)))
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

	fPurchaseOrderDocument, _ := c.FormFile("document")
	if fPurchaseOrderDocument == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field document is required"), "cannot submit if document is empty"))
		return
	}

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

	if err := utils.EmailFormatValidation(p.RecipientEmail); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, "Invalid email recipient format"))
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
		Document:            fPurchaseOrderDocument,
		OperatingActivityID: operatingActivityID,
	}

	if PurchaseOrder, err := services.Handler.RetrievePurchaseOrder(PurchaseOrderId); err == nil {

		// check existing operating id
		if PurchaseOrder.OperatingActivity.ID.String() != "" {
			if err := services.Handler.CheckExistingPurchaseOrder(id, struct {
				*models.PurchaseOrder
			}{&models.PurchaseOrder{
				OperatingActivityID: repopulateFormDataPayload.OperatingActivityID,
			}}); err != nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
				return
			}
		}

		// check duplication operating id
		if PurchaseOrder.OperatingActivity.ID != repopulateFormDataPayload.OperatingActivityID {
			if err := services.Handler.CheckExistingPurchaseOrder(id, struct {
				*models.PurchaseOrder
			}{&models.PurchaseOrder{
				OperatingActivityID: repopulateFormDataPayload.OperatingActivityID,
			}}); err == nil {
				c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("data cannot be duplicated"), fmt.Sprintf("data is already existing with operating activity's tax number %s", PurchaseOrder.OperatingActivity.TaxInvoiceNumber)))
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
