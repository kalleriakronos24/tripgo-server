package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models/master"
)

type CheckExistingEntityCustomerStruct struct {
	*master.Customer
}

func (module *module) RetrieveEntityCustomerByUserID(userId uuid.UUID) (m master.Customer, err error) {
	if m, err = module.db.userCustomerModel.GetOneByID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveAllEntityCustomer(id uuid.UUID) (m []master.Customer, err error) {
	return
}

// func (module *module) InsertEntityCustomer(p *dto.InsertEntityCustomer) (err error) {

// 	tx := database.GetDatabaseConnection().Begin()

// 	// get all available Customers

// 	if err = module.db.EntityCustomer.InsertEntityCustomer(models.EntityCustomer{
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

// func (module *module) UpdateEntityCustomer(id uuid.UUID, p *dto.UpdateEntityCustomer) (err error) {
// 	tx := database.GetDatabaseConnection().Begin()
// 	if err = module.db.EntityCustomer.UpdateEntityCustomer(id, models.EntityCustomer{
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

func (module *module) RetrieveAllEntityCustomerPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveEntityCustomer(id uuid.UUID) (m master.Customer, err error) {
	return
}
