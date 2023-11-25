package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
)

func (module *module) RetrieveUser(id uuid.UUID) (m masterModels.User, err error) {
	if m, err = module.db.userModel.GetOneByID(id); err != nil {
		return masterModels.User{}, fmt.Errorf("user not found")
	}
	return
}

func (module *module) UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error) {
	if err = module.db.userModel.UpdateUser(id, masterModels.User{
		ID:    id,
		Email: p.Email,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}
