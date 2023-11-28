package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
)

type CheckExistingPaymentStruct struct {
	*models.Payment
}

func (module *module) RetrieveAllPayment(id uuid.UUID) (m []models.Payment, err error) {
	if m, err = module.db.paymentModel.GetAllPayment(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrievePayment(id uuid.UUID) (m models.Payment, err error) {
	if m, err = module.db.paymentModel.GetOnePaymentByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertPayment(p *dto.InsertPayment) (err error) {
	if err = module.db.paymentModel.InsertPayment(models.Payment{
		Currency:            p.Currency,
		Amount:              p.Amount,
		Term:                p.Term,
		ARAging:             p.ARAging,
		Remarks:             p.Remarks,
		DueDate:             p.DueDate,
		SettlementDate:      p.SettlementDate,
		OperatingActivityID: p.OperatingActivityID,
		PaymentCreatedBy:    p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdatePayment(id uuid.UUID, p *dto.UpdatePayment) (err error) {
	if err = module.db.paymentModel.UpdatePayment(id, models.Payment{
		Currency:            p.Currency,
		Amount:              p.Amount,
		Term:                p.Term,
		ARAging:             p.ARAging,
		Remarks:             p.Remarks,
		DueDate:             p.DueDate,
		SettlementDate:      p.SettlementDate,
		OperatingActivityID: p.OperatingActivityID,
		PaymentUpdatedBy:    p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingPayment(id string, param CheckExistingPaymentStruct) (err error) {

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.paymentModel.GetOnePaymentByOperatingActivityID(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.paymentModel.GetOnePaymentByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
