package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/utils"
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

func (module *module) RetrieveBookingTransferAcceptedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedAcceptedByDriverID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveBookingTransferCancelledByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedCancelledByDriverID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveBookingTransferOngoingByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedOnGoingByDriverID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveBookingTransferCompletedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedCompletedByDriverID(userId); err != nil {
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
		Status: "driver accepted, please wait for pickup",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewTrue(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Booking Transfer Accepted",
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
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewTrue(),
		IsOnGoing:   utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Booking Transfer Cancelled",
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

func (module *module) OngoingBookingTransfer(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "driver on it's way to pickup location",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewTrue(),
		IsCompleted: utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Driver on it's way to pickup location",
		Body: `<html><body>
		<p>Sit tight, wait driver arrived at your pickup location and be ready travel to the destination.</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}

	tx.Commit()
	return
}

func (module *module) CompleteBookingTransfer(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "your rides completed. thank you :)",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewFalse(),
		IsCompleted: utils.NewTrue(),
		IsPickedUp:  utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Rides Completed, have a great day :)",
		Body: `<html><body>
		<p>Thank you for using our service</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	tx.Commit()
	return
}

func (module *module) CompletePickupBooking(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "pickup complete, heading to destination",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
		IsPickedUp:  utils.NewTrue(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Pickup Completed, Heading to your Destination",
		Body: `<html><body>
		<p>Thank you for using our service</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	tx.Commit()
	return
}
