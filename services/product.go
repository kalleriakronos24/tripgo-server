package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"strings"
)

type CheckExistingProductStruct struct {
	*models.Product
}

func (module *module) RetrieveAllProduct(id uuid.UUID) (m []models.Product, err error) {
	if m, err = module.db.productModel.GetAllProduct(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllProductPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.productModel.GetAllProductPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveProduct(id uuid.UUID) (m models.Product, err error) {
	if m, err = module.db.productModel.GetOneProductByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertProduct(p *dto.InsertProduct) (err error) {
	if err = module.db.productModel.InsertProduct(models.Product{
		Name:             strings.ToUpper(p.Name),
		UnitPrice:        p.UnitPrice,
		Packaging:        p.Packaging,
		Stock:            p.Stock,
		Note:             p.Note,
		CompanyID:        p.CompanyID,
		VATIncluded:      p.VATIncluded,
		ProductCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateProduct(id uuid.UUID, p *dto.UpdateProduct) (err error) {
	if err = module.db.productModel.UpdateProduct(id, models.Product{
		Name:             strings.ToUpper(p.Name),
		UnitPrice:        p.UnitPrice,
		Packaging:        p.Packaging,
		Stock:            p.Stock,
		Note:             p.Note,
		CompanyID:        p.CompanyID,
		VATIncluded:      p.VATIncluded,
		ProductUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) DeleteProduct(id uuid.UUID) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	if err = module.db.productModel.DeleteProduct(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.operatingActivityProductModel.DeleteOperatingActivityProductByProductID(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.productHistoryModel.DeleteProductHistoryByProductID(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.purchaseOrderProductModel.DeletePurchaseOrderProductByProductID(id, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()

	return
}

func (module *module) CheckExistingProduct(id string, param CheckExistingProductStruct) (err error) {

	if param.Name != "" {
		if _, dbErr := module.db.productModel.GetOneProductByName(param.Name); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.productModel.GetOneProductByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
