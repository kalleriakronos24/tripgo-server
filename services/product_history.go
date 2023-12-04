package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
)

type CheckExistingProductHistoryStruct struct {
	*models.ProductHistory
}

func (module *module) RetrieveAllProductHistory(id uuid.UUID) (m []models.ProductHistory, err error) {
	if m, err = module.db.productHistoryModel.GetAllProductHistory(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllProductHistoryPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.productHistoryModel.GetAllProductHistoryPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveProductHistory(id uuid.UUID) (m models.ProductHistory, err error) {
	if m, err = module.db.productHistoryModel.GetOneProductHistoryByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertProductHistory(p *dto.InsertProductHistory) (err error) {
	if err = module.db.productHistoryModel.InsertProductHistory(models.ProductHistory{
		Status:                  p.Status,
		Quantity:                p.Quantity,
		ProductID:               p.ProductID,
		OperatingActivityID:     p.OperatingActivityID,
		ProductHistoryCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateProductHistory(id uuid.UUID, p *dto.UpdateProductHistory) (err error) {
	if err = module.db.productHistoryModel.UpdateProductHistory(id, models.ProductHistory{
		Status:                  p.Status,
		Quantity:                p.Quantity,
		ProductID:               p.ProductID,
		OperatingActivityID:     p.OperatingActivityID,
		ProductHistoryUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingProductHistory(id string, param CheckExistingProductHistoryStruct) (err error) {

	if param.ProductID.String() != "" {
		if _, dbErr := module.db.productHistoryModel.GetOneProductHistoryByProductID(param.ProductID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.productHistoryModel.GetOneProductHistoryByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
