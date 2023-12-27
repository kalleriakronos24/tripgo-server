package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
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

	return
}

func (module *module) DeleteCompany(id uuid.UUID, userId uuid.UUID) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	if err = module.db.companyModel.DeleteCompany(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.userModel.UpdateUserToInactive(userId, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.userModel.RemoveUserFromCompany(userId, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.productModel.DeleteProductByCompanyID(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	var clients []masterModels.Client
	var clientErr error
	if clients, clientErr = module.db.clientModel.GetOneClientByCompanyID(id); clientErr != nil {
		//tx.Rollback()
		//return errors.New(clientErr.Error())
	}

	if len(clients) > 0 {

		for _, client := range clients {

			if err = module.db.clientModel.DeleteClientByCompanyID(client.ID, tx); err != nil {
				tx.Rollback()
				return errors.New(err.Error())
			}

			if err = module.db.purchaseOrderProductModel.DeletePurchaseOrderProductByClientID(client.ID, tx); err != nil {
				tx.Rollback()
				return errors.New(err.Error())
			}

			var operatingActivities []models.OperatingActivity
			var optActErr error
			if operatingActivities, optActErr = module.db.operatingActivityModel.GetOneOperatingActivityByClientID(client.ID); optActErr != nil {
				//tx.Rollback()
				//return errors.New(optActErr.Error())
			}

			if len(operatingActivities) > 0 {

				for _, operatingActivity := range operatingActivities {

					if err = module.db.operatingActivityProductModel.DeleteOptActProductByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.deliveryOrderModel.DeleteDeliveryOrderByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.invoiceModel.DeleteInvoiceByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.quotationModel.DeleteQuotationByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.purchaseOrderModel.DeletePurchaseOrderByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.paymentModel.DeletePaymentByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}

					if err = module.db.paymentInstallmentModel.DeletePaymentInstallmentByOptActID(operatingActivity.ID, tx); err != nil {
						tx.Rollback()
						return errors.New(err.Error())
					}
				}
			}
		}
	}

	tx.Commit()
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
