package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/booklap-be/db"
	masterModels "github.com/kalleriakronos24/booklap-be/models/master"
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
