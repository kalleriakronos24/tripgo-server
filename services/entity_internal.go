package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models/master"
)

type CheckExistingEntityInternalStruct struct {
	*master.Driver
}

func (module *module) RetrieveEntityInternalByUserID(userId uuid.UUID) (m master.Internal, err error) {
	if m, err = module.db.userInternalModel.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get internal user information")
	}
	return
}

func (module *module) RetrieveAllEntityInternal(id uuid.UUID) (m []master.Internal, err error) {
	return
}

func (module *module) RetrieveAllEntityInternalPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveEntityInternal(id uuid.UUID) (m master.Driver, err error) {
	return
}
