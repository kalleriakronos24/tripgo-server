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
	"os"
)

func GenerateSPHDocument(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	invoiceId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if Data, err := services.Handler.GenerateSPHDocument(invoiceId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, ""))
		return
	} else {
		fileName := fmt.Sprintf("attachment; filename=%s.pdf", Data.FileName)
		byteFile, err := os.ReadFile(Data.OutputPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("failed to retrieve sph pdf"), ""))
			return
		}
		c.Header("Content-Disposition", fileName)
		c.Data(http.StatusOK, "application/pdf", byteFile)
		err = os.Remove(Data.OutputPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("failed to remove sph pdf"), ""))
			return
		}
		return
	}
}

func GETAllQuotation(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if quotation, err := services.Handler.RetrieveAllQuotationPaginated(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	} else {
		c.JSON(http.StatusOK, dto.Response{Data: quotation})
	}

}

func GETQuotation(c *gin.Context) {
	var err error

	id, _ := c.Params.Get("id")
	quotationId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	var quotation models.Quotation
	if quotation, err = services.Handler.RetrieveQuotation(quotationId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: quotation})
}

func POSTQuotation(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.InsertQuotationValidator{CreatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.InsertQuotation
	if pValidator.OperatingActivityID != "" {
		operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)

		p = &dto.InsertQuotation{
			FrancoArea:          pValidator.FrancoArea,
			PaymentTerm:         pValidator.PaymentTerm,
			SendAfter:           pValidator.SendAfter,
			Date:                utils.ConvertStrToDateTime(pValidator.Date),
			OperatingActivityID: operatingActivityId,
			CreatedBy:           pValidator.CreatedBy,
		}
	} else {

		p = &dto.InsertQuotation{
			FrancoArea:  pValidator.FrancoArea,
			PaymentTerm: pValidator.PaymentTerm,
			SendAfter:   pValidator.SendAfter,
			Date:        utils.ConvertStrToDateTime(pValidator.Date),
			CreatedBy:   pValidator.CreatedBy,
		}
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if err = services.Handler.InsertQuotation(p); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "quotation"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}

func PUTQuotation(c *gin.Context) {
	var err error
	userLoggedInId := c.GetString("user_id")
	userId, err := uuid.Parse(userLoggedInId)
	pValidator := &dto.UpdateQuotationValidator{UpdatedBy: userId}

	if err = c.ShouldBindJSON(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var p *dto.UpdateQuotation
	if pValidator.OperatingActivityID != "" {
		operatingActivityId, _ := uuid.Parse(pValidator.OperatingActivityID)

		p = &dto.UpdateQuotation{
			FrancoArea:          pValidator.FrancoArea,
			PaymentTerm:         pValidator.PaymentTerm,
			SendAfter:           pValidator.SendAfter,
			Date:                utils.ConvertStrToDateTime(pValidator.Date),
			OperatingActivityID: operatingActivityId,
			UpdatedBy:           pValidator.UpdatedBy,
		}
	} else {
		p = &dto.UpdateQuotation{
			FrancoArea:  pValidator.FrancoArea,
			PaymentTerm: pValidator.PaymentTerm,
			SendAfter:   pValidator.SendAfter,
			Date:        utils.ConvertStrToDateTime(pValidator.Date),
			UpdatedBy:   pValidator.UpdatedBy,
		}
	}

	id, _ := c.Params.Get("id")
	quotationId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("uuid-error", err, ""))
		return
	}

	if err := services.Handler.CheckExistingOperatingActivity(p.OperatingActivityID.String(), struct{ *models.OperatingActivity }{&models.OperatingActivity{}}); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", err, fmt.Sprintf("operating id %s is not found", p.OperatingActivityID)))
		return
	}

	if _, err := services.Handler.RetrieveQuotation(quotationId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "quotation"))
		return
	}

	if err = services.Handler.UpdateQuotation(quotationId, p); err != nil {
		c.JSON(http.StatusNotModified, constants.GetErrorResponse("update-failed", err, "quotation"))
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "success"})
}
