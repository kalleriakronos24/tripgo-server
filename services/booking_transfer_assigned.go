package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
)

type CheckExistingBookingTransferAssignedStruct struct {
	*models.BookingTransfer
}

func (module *module) RetrieveBookingTransferAssignedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedNotAcceptedByDriverID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) AcceptBookingTransfer(id uuid.UUID) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "driver will pickup you before pickup date",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted: true,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "notification@wadahgo.com",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "WadahGo - Booking Transfer Accepted",
		Body: `<html><body>
		<p>Driver Accepted your booking, please wait driver pick you up at the pick up date scheduled.</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	tx.Commit()
	return
}

func (module *module) CancelBookingTransfer(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "driver cancelled due some reasons",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  false,
		IsCancelled: true,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "notification@wadahgo.com",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "WadahGo - Booking Transfer Cancelled",
		Body: `<html><body>
		<p>Oops.. Driver canceled your booking due to some reason :( please create a new booking</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}

	// logic to send email to customer
	tx.Commit()
	return
}
