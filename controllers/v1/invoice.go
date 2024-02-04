package v1

import (
	"errors"
	"fmt"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/services"
)

func GenerateInvoiceDocument(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	invoiceId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if Data, err := services.Handler.GenerateInvoiceDocument(invoiceId); err != nil {
		log.Printf("error >>> %s", err)
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, ""))
		return
	} else {

		log.Printf("%v", Data)
		fileName := fmt.Sprintf("attachment; filename=%s.pdf", Data.FileName)
		byteFile, err := os.ReadFile(Data.OutputPath)
		if err != nil {
			log.Printf("error >>> %s", err.Error())
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("failed to retrieve invoice pdf"), ""))
			return
		}
		c.Header("Content-Disposition", fileName)
		c.Data(http.StatusOK, "application/pdf", byteFile)
		err = os.Remove(Data.OutputPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("failed to remove invoice pdf"), ""))
			return
		}
		return
	}
}

func GETAllInvoice(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if Invoice, err := services.Handler.RetrieveAllInvoicePaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "invoice"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: Invoice})
	}

}

func GETInvoice(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	InvoiceId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var Invoice models.Invoice
	if Invoice, err = services.Handler.RetrieveInvoice(InvoiceId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "invoice"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: Invoice})
}

func POSTInvoice(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertInvoiceValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.InsertInvoice{
		Type:                pValidator.Type,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		OperatingActivityID: operatingActivityId,
		CreatedBy:           pValidator.CreatedBy,
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err = services.Handler.InsertInvoice(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "invoice"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTInvoice(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateInvoiceValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)
	p := &dto.UpdateInvoice{
		Type:                pValidator.Type,
		Date:                utils.ConvertStrToDateTime(pValidator.Date),
		OperatingActivityID: operatingActivityId,
		UpdatedBy:           pValidator.UpdatedBy,
	}

	id, _ := c.Params.Get("id")
	InvoiceId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if _, err := services.Handler.RetrieveInvoice(InvoiceId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "invoice"))
		return
	}

	if err = services.Handler.UpdateInvoice(InvoiceId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "invoice"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
