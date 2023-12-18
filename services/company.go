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

type CheckExistingCompanyStruct struct {
	*masterModels.Company
}

func (module *module) RetrieveCompanyByUserID(userId uuid.UUID) (m masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetOneCompanyByUserID(userId); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllCompany(id uuid.UUID) (m []masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetAllCompany(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllCompanyPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.companyModel.GetAllCompanyPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveCompany(id uuid.UUID) (m masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetOneCompanyByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertCompany(p *dto.InsertCompany) (err error) {
	tx := database.GetDatabaseConnection()
	companyModel := masterModels.Company{
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
	}

	if errCompany := tx.Create(&companyModel); err != nil {
		tx.Rollback()
		return errCompany.Error
	}

	userModel := masterModels.User{
		ID:        p.CreatedBy,
		CompanyID: companyModel.ID,
	}

	if errUser := tx.Updates(&userModel); err != nil {
		tx.Rollback()
		return errUser.Error
	}

	tx.Commit()
	con, _ := database.GetDatabaseConnection().DB()
	_ = con.Close()
	return
}

func (module *module) UpdateCompany(id uuid.UUID, p *dto.UpdateCompany) (err error) {

	tx := database.GetDatabaseConnection()
	companyModel := masterModels.Company{
		ID:                id,
		Name:              p.Name,
		PhoneNumber:       p.PhoneNumber,
		Email:             strings.ToLower(p.Email),
		Address:           p.Address,
		PICName:           p.PICName,
		PICDesignation:    p.PICDesignation,
		BankAccountName:   p.BankAccountName,
		BankAccountNumber: p.BankAccountNumber,
		CompanyCreatedBy:  p.UpdatedBy,
		CompanyUpdatedBy:  p.UpdatedBy,
	}

	if errCompany := tx.Updates(&companyModel); err != nil {
		tx.Rollback()
		return errCompany.Error
	}

	userModel := masterModels.User{
		ID:        p.UpdatedBy,
		CompanyID: companyModel.ID,
	}

	if errUser := tx.Updates(&userModel); err != nil {
		tx.Rollback()
		return errUser.Error
	}

	tx.Commit()
	con, _ := database.GetDatabaseConnection().DB()
	_ = con.Close()

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
