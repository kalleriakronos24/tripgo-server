package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
)

func (module *module) RetrieveUser(username string) (m masterModels.User, err error) {
	if m, err = module.db.userModel.GetOneByUsername(username); err != nil {
		return masterModels.User{}, fmt.Errorf("cannot find user with username %s", username)
	}
	return
}

func (module *module) UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error) {
	if err = module.db.userModel.UpdateUser(masterModels.User{
		ID:    id,
		Email: p.Email,
		Bio:   p.Bio,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}
