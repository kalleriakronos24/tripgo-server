package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/onesignal"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

type CheckExistingBookingTransferStruct struct {
	*models.BookingTransfer
}

func (module *module) RetrieveBookingTransferByUserID(userId uuid.UUID) (m models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetOneByID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveAllBookingTransferByCustomer(id uuid.UUID) (m []*models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetAllByCustomerID(id); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) InsertBookingTransfer(p *dto.InsertBookingTransfer) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	// get all available drivers
	// with
	// 1. enough balance
	// 2. driver is active and not on ride/active booking
	// 3. has the request car from the customer
	// 4. internal driver has 90% chance to get order
	// 5. external driver has 20% chance to get order
	var driverBalance *models.BalanceDriver
	if driverBalance, err = module.db.balanceDriver.GetAllDriverHasEnoughBalance(float64(p.Price), p.CarModelID); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	var carManagement master.CarManagement
	if carManagement, err = module.db.carManagementModel.GetOneByCarModelIDAndAvailable(p.CarModelID, driverBalance.DriverID); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	var bookingTransfer models.BookingTransfer
	if bookingTransfer, err = module.db.bookingTransfer.InsertBookingTransfer(models.BookingTransfer{
		AdultSeater:       p.AdultSeater,
		ChildSeater:       p.ChildSeater,
		FromLatCoordinate: p.FromLatCoordinate,
		FromLngCoordinate: p.FromLngCoordinate,
		FromLocation:      p.FromLocation,
		ToLocation:        p.ToLocation,
		ToLatCoordinate:   p.ToLatCoordinate,
		ToLngCoordinate:   p.ToLngCoordinate,
		Distance:          p.TotalDistance,
		PassengerNotes:    p.PassengerNotes,
		PickUpDate:        p.PickUpDate,
		Price:             p.Price,
		CarModelID:        p.CarModelID,
		CustomerID:        p.CustomerID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// immediately assign to the driver
	if err = module.db.bookingTransferAssigned.InsertBookingTransferAssigned(models.BookingTransferAssigned{
		CarManagementID:   carManagement.ID,
		BookingTransferID: bookingTransfer.ID,
		IsAccepted:        utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	formattedNotificationMessage := fmt.Sprintf("Booking Transfer Request \n %v \n %v Person Pax x %v Luggagge \n Pickup Date %v \n Notes: %v \n Price: %v", carManagement.Name, carManagement.CarModel.PersonCount, carManagement.CarModel.LuggageCount, utils.ConvertEnToIDDateTime(bookingTransfer.PickUpDate), bookingTransfer.PassengerNotes, bookingTransfer.Price)

	// send notification to the selected driver
	onesignal.PushNotificationSingleExternalId(driverBalance.Driver.Credentials.Email, formattedNotificationMessage)

	tx.Commit()
	return
}

func (module *module) UpdateBookingTransfer(id uuid.UUID, p *dto.UpdateBookingTransfer) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	if err = module.db.bookingTransfer.UpdateBookingTransfer(id, models.BookingTransfer{
		AdultSeater:       p.AdultSeater,
		ChildSeater:       p.ChildSeater,
		FromLatCoordinate: p.FromLatCoordinate,
		FromLngCoordinate: p.FromLngCoordinate,
		FromLocation:      p.FromLocation,
		ToLocation:        p.ToLocation,
		ToLatCoordinate:   p.ToLatCoordinate,
		ToLngCoordinate:   p.ToLngCoordinate,
		PassengerNotes:    p.PassengerNotes,
		PickUpDate:        p.PickUpDate,
		Price:             p.Price,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// update driver status to busy and put the car busy also, and waiting to approval by the driver
	tx.Commit()
	return
}

func (module *module) RetrieveAllBookingTransferPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveBookingTransfer(id uuid.UUID) (m models.BookingTransfer, err error) {
	return
}
