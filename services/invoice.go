package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
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

func (module *module) InsertInvoice(p *dto.InsertInvoice) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	Invoice := models.Invoice{
		Number:              p.Number,
		Type:                p.Type,
		Date:                p.Date,
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
		Number:              p.Number,
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
