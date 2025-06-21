package services

import (
	"errors"

	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models"
)

type CheckExistingPaymentStruct struct {
	*models.Payment
}

func (module *module) RetrievePaymentByUserID(userId uuid.UUID) (m models.Payment, err error) {
	if m, err = module.db.paymentModel.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get booking transfers rating")
	}
	return
}

func (module *module) InsertPayment(p *models.Payment) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	// immediately assign to the driver
	if err = module.db.paymentModel.InsertPayment(models.Payment{
		Amount:            p.Amount,
		PaymentMethod:     p.PaymentMethod,
		Status:            p.Status,
		CustomerID:        p.CustomerID,
		BookingTransferID: p.BookingTransferID,
		DriverID:          p.DriverID,
		PI:                p.PI,
		Currency:          p.Currency,
		CardLastNumber:    p.CardLastNumber,
		ReceiptURL:        p.ReceiptURL,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	return err
}
