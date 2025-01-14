package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

type CheckExistingBookingTransferAssignedStruct struct {
	*models.BookingTransfer
}

func (module *module) RetrieveBookingTransferAssignedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedNotAcceptedByDriverID(userId); err != nil {
		return m, errors.New("failed to get assigned bookings")
	}
	return m, nil
}

func (module *module) RetrieveBookingTransferAcceptedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedAcceptedByDriverID(userId); err != nil {
		return m, errors.New("failed to get accepted bookings")
	}
	return m, nil
}

func (module *module) RetrieveBookingTransferCancelledByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedCancelledByDriverID(userId); err != nil {
		return m, errors.New("failed to get cancelled bookings")
	}
	return m, nil
}

func (module *module) RetrieveBookingTransferOngoingByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedOnGoingByDriverID(userId); err != nil {
		return m, errors.New("failed to get ongoing bookings")
	}
	return m, nil
}

func (module *module) RetrieveBookingTransferCompletedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error) {
	if m, err = module.db.bookingTransferAssigned.GetBookingAssignedCompletedByDriverID(userId); err != nil {
		return m, errors.New("failed to get completed bookings")
	}
	return m, nil
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

	// once accepted the order immediately deduct the driver balance
	var balanceDriver models.BalanceDriver
	if balanceDriver, err = module.db.balanceDriver.GetOneByID(bookingTransferAssigned.DriverID); err != nil {
		return errors.New(err.Error())
	}

	if BalanceDriverErr := tx.Model(&models.BalanceDriver{}).Where("driver_id", bookingTransferAssigned.DriverID).Updates(&models.BalanceDriver{
		Amount: balanceDriver.Amount - float64(bookingTransferAssigned.BookingTransfer.Price)*0.14,
	}); BalanceDriverErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to deduct driver balance")
	}

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	month := int(currentMonth)
	randomUid, _ := utils.GenerateNumber(10)

	// config.AppConfig.APPUrl
	DriverTopup := models.DriverTopup{
		Uid:      fmt.Sprintf("DRV/DT/%v%v/%v", utils.IntegerToRoman(currentYear), utils.IntegerToRoman(month), randomUid),
		Amount:   -float64(bookingTransferAssigned.BookingTransfer.Price) * 0.14,
		DriverID: bookingTransferAssigned.DriverID,
	}

	if DriverTopupErr := tx.Create(&DriverTopup); DriverTopupErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload balance topup")
	}

	DriverTransactionHistory := models.DriverTransactionHistory{
		Remark:        "order deduction",
		Status:        "deduction",
		DriverID:      bookingTransferAssigned.DriverID,
		DriverTopupID: DriverTopup.ID,
	}

	if DriverTransactionHistoryErr := tx.Create(&DriverTransactionHistory); DriverTransactionHistoryErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload balance topup")
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
		return errors.New("failed to cancel booking")
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewTrue(),
		IsOnGoing:   utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to cancel booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Booking Transfer Cancelled",
		Body: `<html><body>
		<p>Oops.. Driver canceled your booking due to some reason :( please create a new booking</p>
		</body></html>`,
	}); err != nil {
		return errors.New("server error. please try again later")
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
		return errors.New("failed to set ongoing booking")
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewTrue(),
		IsCompleted: utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to set ongoing booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Driver on it's way to pickup location",
		Body: `<html><body>
		<p>Sit tight, wait driver arrived at your pickup location and be ready travel to the destination.</p>
		</body></html>`,
	}); err != nil {
		return errors.New("server error. please try again later")
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
		return errors.New("failed to set complete booking")
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewFalse(),
		IsCompleted: utils.NewTrue(),
		IsPickedUp:  utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to set complete booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Rides Completed, have a great day :)",
		Body: `<html><body>
		<p>Thank you for using our service</p>
		</body></html>`,
	}); err != nil {
		return errors.New("server error. please try again later")
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
		return errors.New("failed to set pickup complete booking")
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewFalse(),
		IsOnGoing:   utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
		IsPickedUp:  utils.NewTrue(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to set pickup complete booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: "Pickup Completed, Heading to your Destination",
		Body: `<html><body>
		<p>Thank you for using our service</p>
		</body></html>`,
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) SwitchDriver(id uuid.UUID, driverId uuid.UUID, plateNumber string) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	var carManagement master.CarManagement
	if carManagement, err = module.db.carManagementModel.GetSameCarManagementByAgentByPlateNumberAndCompanyIdAndDriverId(bookingTransferAssigned.Driver.Company.ID, plateNumber, driverId); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		DriverID:        driverId,
		CarManagementID: carManagement.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to set pickup complete booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.BookingTransfer.Customer.Credentials.Email,
		Subject: fmt.Sprintf("Your BookingID: %v, driver has changed", bookingTransferAssigned.BookingTransfer.Uid),
		Body: fmt.Sprintf(`<html><body>
		<p>Your new driver for BookingID: %v</p>
		<p>Driver Name: %v</p>
		<p>Plate Number: %v</p>
		</body></html>`, bookingTransferAssigned.BookingTransfer.Uid, carManagement.Driver.Name, carManagement.PlateNumber),
	}); err != nil {
		return errors.New("server error. please try again later")
	}

	tx.Commit()
	return err
}
