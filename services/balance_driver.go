package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models"
)

func (module *module) RetrieveDriverBalanceDetailByDriver(id uuid.UUID) (m models.BalanceDriver, err error) {
	if m, err = module.db.balanceDriver.GetOneByID(id); err != nil {
		return m, errors.New("failed to get wallet information")
	}
	return
}
