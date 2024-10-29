package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models/master"
)

type CheckExistingEntityDriverStruct struct {
	*master.Driver
}

func (module *module) RetrieveEntityDriverByUserID(userId uuid.UUID) (m master.Driver, err error) {
	if m, err = module.db.driverModel.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get user information")
	}
	return
}

func (module *module) RetrieveAllEntityDriver(id uuid.UUID) (m []master.Driver, err error) {
	return
}

// func (module *module) InsertEntityDriver(p *dto.InsertEntityDriver) (err error) {

// 	tx := database.GetDatabaseConnection().Begin()

// 	// get all available drivers

// 	if err = module.db.EntityDriver.InsertEntityDriver(models.EntityDriver{
// 		AdultSeater:       p.AdultSeater,
// 		ChildSeater:       p.ChildSeater,
// 		FromLatCoordinate: p.FromLatCoordinate,
// 		FromLngCoordinate: p.FromLngCoordinate,
// 		FromLocation:      p.FromLocation,
// 		ToLocation:        p.ToLocation,
// 		ToLatCoordinate:   p.ToLatCoordinate,
// 		ToLngCoordinate:   p.ToLngCoordinate,
// 		PassengerNotes:    p.PassengerNotes,
// 		PickUpDate:        p.PickUpDate,
// 		Price:             p.Price,
// 	}, tx); err != nil {
// 		tx.Rollback()
// 		return errors.New(err.Error())
// 	}

// 	tx.Commit()
// 	return
// }

// func (module *module) UpdateEntityDriver(id uuid.UUID, p *dto.UpdateEntityDriver) (err error) {
// 	tx := database.GetDatabaseConnection().Begin()
// 	if err = module.db.EntityDriver.UpdateEntityDriver(id, models.EntityDriver{
// 		AdultSeater:       p.AdultSeater,
// 		ChildSeater:       p.ChildSeater,
// 		FromLatCoordinate: p.FromLatCoordinate,
// 		FromLngCoordinate: p.FromLngCoordinate,
// 		FromLocation:      p.FromLocation,
// 		ToLocation:        p.ToLocation,
// 		ToLatCoordinate:   p.ToLatCoordinate,
// 		ToLngCoordinate:   p.ToLngCoordinate,
// 		PassengerNotes:    p.PassengerNotes,
// 		PickUpDate:        p.PickUpDate,
// 		Price:             p.Price,
// 	}, tx); err != nil {
// 		tx.Rollback()
// 		return errors.New(err.Error())
// 	}
// 	tx.Commit()
// 	return
// }

func (module *module) RetrieveAllEntityDriverPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveEntityDriver(id uuid.UUID) (m master.Driver, err error) {
	return
}
