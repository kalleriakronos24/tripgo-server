package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
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
