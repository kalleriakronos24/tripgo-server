package services

import (
	"errors"

	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
)

type CheckExistingBookingTransferRatingStruct struct {
	*models.BookingTransferRating
}

func (module *module) RetrieveBookingTransferRatingByUserID(userId uuid.UUID) (m models.BookingTransferRating, err error) {
	if m, err = module.db.bookingTransferRating.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get booking transfers rating")
	}
	return
}

func (module *module) InsertBookingTransferRating(p *dto.InsertBookingTransferRating) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	// immediately assign to the driver
	if err = module.db.bookingTransferRating.InsertBookingTransferRating(models.BookingTransferRating{
		Rating:                    p.Rating,
		CarManagementID:           p.CarManagementId,
		BookingTransferAssignedID: p.BookingTransferAssignedID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	return err
}
