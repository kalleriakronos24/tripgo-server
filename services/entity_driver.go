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

func (module *module) RetrieveAllEntityDriverPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveEntityDriver(id uuid.UUID) (m master.Driver, err error) {
	return
}
