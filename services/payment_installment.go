package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
)

type CheckExistingPaymentInstallmentStruct struct {
	*models.PaymentInstallment
}

func (module *module) RetrieveAllPaymentInstallment(id uuid.UUID) (m []models.PaymentInstallment, err error) {
	if m, err = module.db.paymentInstallmentModel.GetAllPaymentInstallment(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllPaymentInstallmentPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.paymentInstallmentModel.GetAllPaymentInstallmentPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrievePaymentInstallment(id uuid.UUID) (m models.PaymentInstallment, err error) {
	if m, err = module.db.paymentInstallmentModel.GetOnePaymentInstallmentByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertPaymentInstallment(p *dto.InsertPaymentInstallment) (err error) {
	if err = module.db.paymentInstallmentModel.InsertPaymentInstallment(models.PaymentInstallment{
		Currency:                    p.Currency,
		Amount:                      p.Amount,
		Note:                        p.Note,
		Date:                        p.Date,
		PaymentID:                   p.PaymentID,
		PaymentInstallmentCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdatePaymentInstallment(id uuid.UUID, p *dto.UpdatePaymentInstallment) (err error) {
	if err = module.db.paymentInstallmentModel.UpdatePaymentInstallment(id, models.PaymentInstallment{
		Currency:                    p.Currency,
		Amount:                      p.Amount,
		Note:                        p.Note,
		Date:                        p.Date,
		PaymentID:                   p.PaymentID,
		PaymentInstallmentUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingPaymentInstallment(id string, param CheckExistingPaymentInstallmentStruct) (err error) {

	if param.PaymentID.String() != "" {
		if _, dbErr := module.db.paymentInstallmentModel.GetOnePaymentInstallmentByPaymentID(param.PaymentID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.paymentInstallmentModel.GetOnePaymentInstallmentByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
