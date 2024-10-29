package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
)

type CheckExistingUserStruct struct {
	*masterModels.Credentials
}

func (module *module) RetrieveUser(id uuid.UUID) (m masterModels.User, err error) {
	if m, err = module.db.userModel.GetOneByID(id); err != nil {
		return masterModels.User{}, fmt.Errorf("user not found")
	}
	return
}

func (module *module) RetrieveAllUserPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.userModel.GetAllUserPaginated(c, id); err != nil {
		return pagination, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) DeleteUser(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	if err = module.db.userModel.DeleteUser(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}
	tx.Commit()

	return
}

func (module *module) CheckExistingUser(id string, param CheckExistingUserStruct) (err error) {
	if param.Email != "" {
		if _, dbErr := module.db.credentialModel.GetOneByEmail(param.Email); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.credentialModel.GetOneByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
