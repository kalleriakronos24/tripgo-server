package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gorm.io/gorm"
)

type CheckExistingOperatingActivityProductStruct struct {
	*models.OperatingActivityProduct
}

func (module *module) RetrieveAllOperatingActivityProduct(id uuid.UUID) (m []models.OperatingActivityProduct, err error) {
	if m, err = module.db.operatingActivityProductModel.GetAllOperatingActivityProduct(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllOperatingActivityProductPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.operatingActivityProductModel.GetAllOperatingActivityProductPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveOperatingActivityProduct(id uuid.UUID) (m models.OperatingActivityProduct, err error) {
	if m, err = module.db.operatingActivityProductModel.GetOneOperatingActivityProductByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveManyOperatingActivityProductByOperatingActivityID(id uuid.UUID) (m []models.OperatingActivityProduct, err error) {
	if m, err = module.db.operatingActivityProductModel.GetManyOperatingActivityProductByOperatingActivityId(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertOperatingActivityProduct(p *dto.InsertOperatingActivityProduct) (err error) {

	if product, err := module.db.productModel.GetOneProductByID(p.ProductID); err != nil {
		return errors.New(err.Error())
	} else {
		if p.Quantity > product.Stock {
			return errors.New("operating product quantity cannot be more than product stock")
		}
	}

	if err = module.db.operatingActivityProductModel.InsertOperatingActivityProduct(models.OperatingActivityProduct{
		Quantity:                          p.Quantity,
		VATRate:                           p.VATRate,
		SubTotal:                          p.SubTotal,
		GrandTotal:                        p.GrandTotal,
		OperatingActivityID:               p.OperatingActivityID,
		ProductID:                         p.ProductID,
		OperatingActivityProductCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateOperatingActivityProduct(id uuid.UUID, p *dto.UpdateOperatingActivityProduct) (err error) {

	if product, err := module.db.productModel.GetOneProductByID(p.ProductID); err != nil {
		return errors.New(err.Error())
	} else {
		if p.Quantity > product.Stock {
			return errors.New("operating product quantity cannot be more than product stock")
		}
	}

	if err = module.db.operatingActivityProductModel.UpdateOperatingActivityProduct(id, models.OperatingActivityProduct{
		Quantity:                          p.Quantity,
		VATRate:                           p.VATRate,
		SubTotal:                          p.SubTotal,
		GrandTotal:                        p.GrandTotal,
		OperatingActivityID:               p.OperatingActivityID,
		ProductID:                         p.ProductID,
		OperatingActivityProductUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) DeleteOperatingActivityProductByProductID(id uuid.UUID, tx *gorm.DB) (err error) {
	if err = module.db.operatingActivityProductModel.DeleteOperatingActivityProductByProductID(id, tx); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingOperatingActivityProduct(id string, param CheckExistingOperatingActivityProductStruct) (err error) {

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.operatingActivityProductModel.GetOneOperatingActivityProductByOperatingActivityId(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.operatingActivityProductModel.GetOneOperatingActivityProductByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
