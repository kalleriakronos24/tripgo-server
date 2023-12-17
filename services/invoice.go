package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"strconv"
	"strings"
	"time"
)

type CheckExistingInvoiceStruct struct {
	*models.Invoice
}

func (module *module) RetrieveInvoice(id uuid.UUID) (m models.Invoice, err error) {
	if m, err = module.db.invoiceModel.GetOneInvoiceByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllInvoice(id uuid.UUID) (m []models.Invoice, err error) {
	if m, err = module.db.invoiceModel.GetAllInvoice(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllInvoicePaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.invoiceModel.GetAllInvoicePaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertInvoice(p *dto.InsertInvoice) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	InvoiceModel := models.Invoice{}

	invoiceType := strings.ToLower(p.Type)

	_, invoiceLastDataErr := database.GetLastDocumentInvoiceNumber(invoiceType, tx, &InvoiceModel)

	if invoiceLastDataErr != nil {
		InvoiceModel.Sequence = "0001"
	}

	parsedNumber, _ := strconv.Atoi(InvoiceModel.Sequence)
	strLastNumber := strconv.Itoa(parsedNumber)
	getFrontDigitNumber := strings.Replace(InvoiceModel.Sequence, strLastNumber, "", -1)
	convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
	addActiveNumberByOne := convertActiveNumberToStr + 1
	convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
	finalNumber := getFrontDigitNumber + convertAddedActiveNumberToStr

	now := time.Now()
	year := now.Year()
	month := now.Month()
	monthInRoman := utils.IntegerToRoman(int(month))

	var formattedNumber string

	if p.Type == "proforma" {
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "PRF", monthInRoman, year)
	}

	if p.Type == "invoice" {
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", finalNumber, "INV", monthInRoman, year)
	}

	Invoice := models.Invoice{
		Number:              formattedNumber,
		Type:                invoiceType,
		Date:                p.Date,
		Sequence:            finalNumber,
		OperatingActivityID: p.OperatingActivityID,
		InvoiceCreatedBy:    p.CreatedBy,
	}
	if InvoiceErr := tx.Create(&Invoice); err != nil {
		tx.Rollback()
		return InvoiceErr.Error
	}
	tx.Commit()
	return
}

func (module *module) UpdateInvoice(id uuid.UUID, p *dto.UpdateInvoice) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	Invoice := models.Invoice{
		ID:                  id,
		Type:                p.Type,
		Date:                p.Date,
		OperatingActivityID: p.OperatingActivityID,
		InvoiceUpdatedBy:    p.UpdatedBy,
	}
	if InvoiceErr := tx.Updates(&Invoice); err != nil {
		tx.Rollback()
		return InvoiceErr.Error
	}
	tx.Commit()
	return
}

func (module *module) CheckExistingInvoice(id string, param CheckExistingInvoiceStruct) (err error) {

	if param.Number != "" {
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByOperatingActivityID(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.invoiceModel.GetOneInvoiceByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
