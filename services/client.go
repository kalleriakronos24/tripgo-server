package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"strings"
)

type CheckExistingClientStruct struct {
	*masterModels.Client
}

func (module *module) RetrieveAllClient(id uuid.UUID) (m []masterModels.Client, err error) {
	if m, err = module.db.clientModel.GetAllClient(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllClientPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.clientModel.GetAllClientPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveClient(id uuid.UUID) (m masterModels.Client, err error) {
	if m, err = module.db.clientModel.GetOneClientByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertClient(p *dto.InsertClient) (err error) {
	if err = module.db.clientModel.InsertClient(masterModels.Client{
		Name:            p.Name,
		PhoneNumber:     p.PhoneNumber,
		Email:           strings.ToLower(p.Email),
		Address:         p.Address,
		ClientCreatedBy: p.CreatedBy,
		CompanyID:       p.CompanyID,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateClient(id uuid.UUID, p *dto.UpdateClient) (err error) {
	if err = module.db.clientModel.UpdateClient(id, masterModels.Client{
		Name:            p.Name,
		PhoneNumber:     p.PhoneNumber,
		Email:           strings.ToLower(p.Email),
		Address:         p.Address,
		ClientUpdatedBy: p.UpdatedBy,
		CompanyID:       p.CompanyID,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingClient(id string, param CheckExistingClientStruct) (err error) {

	if param.Name != "" {
		if _, dbErr := module.db.clientModel.GetOneClientByName(param.Name); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.Email != "" {
		if _, dbErr := module.db.clientModel.GetOneClientByEmail(param.Email); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.clientModel.GetOneClientByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
