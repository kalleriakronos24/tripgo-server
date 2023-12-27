package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
)

type CheckExistingUserStruct struct {
	*masterModels.User
}

func (module *module) RetrieveUser(id uuid.UUID) (m masterModels.User, err error) {
	if m, err = module.db.userModel.GetOneByID(id); err != nil {
		return masterModels.User{}, fmt.Errorf("user not found")
	}
	return
}

func (module *module) UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error) {

	companyId, _ := uuid.Parse(p.CompanyID)
	if err = module.db.userModel.UpdateUser(id, masterModels.User{
		ID:        id,
		Name:      p.Name,
		Address:   p.Address,
		Username:  p.Username,
		Email:     p.Email,
		Role:      p.Role,
		CompanyID: companyId,
	}); err != nil {
		return errors.New(err.Error())
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

	if param.Name != "" {
		if _, dbErr := module.db.userModel.GetOneByUserName(param.Username); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.Email != "" {
		if _, dbErr := module.db.userModel.GetOneByEmail(param.Email); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.Username != "" {
		if _, dbErr := module.db.userModel.GetOneByUserName(param.Username); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.userModel.GetOneByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
