package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"strings"
)

type CheckExistingCompanyStruct struct {
	*masterModels.Company
}

func (module *module) RetrieveCompany(id uuid.UUID) (m masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetOneCompanyByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertCompany(p *dto.InsertCompany) (err error) {

	//user := masterModels.User{}
	if err = module.db.companyModel.InsertCompany(masterModels.Company{
		Name:              p.Name,
		PhoneNumber:       p.PhoneNumber,
		Email:             strings.ToLower(p.Email),
		Address:           p.Address,
		PICName:           p.PICName,
		PICDesignation:    p.PICDesignation,
		BankAccountName:   p.BankAccountName,
		BankAccountNumber: p.BankAccountNumber,
		CompanyCreatedBy:  p.CreatedBy,
		CompanyUpdatedBy:  p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateCompany(id uuid.UUID, p *dto.UpdateCompany) (err error) {
	if err = module.db.companyModel.UpdateCompany(id, masterModels.Company{
		ID:                p.ID,
		Name:              p.Name,
		PhoneNumber:       p.PhoneNumber,
		Email:             strings.ToLower(p.Email),
		Address:           p.Address,
		PICName:           p.PICName,
		PICDesignation:    p.PICDesignation,
		BankAccountName:   p.BankAccountName,
		BankAccountNumber: p.BankAccountNumber,
		CompanyUpdatedBy:  p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingCompany(id string, param CheckExistingCompanyStruct) (err error) {

	if param.Name != "" {
		if _, dbErr := module.db.companyModel.GetOneCompanyByName(param.Name); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.Email != "" {
		if _, dbErr := module.db.companyModel.GetOneCompanyByEmail(param.Email); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.companyModel.GetOneCompanyByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
