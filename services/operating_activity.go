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

type CheckExistingOperatingActivityStruct struct {
	*models.OperatingActivity
}

func (module *module) RetrieveAllOperatingActivity(id uuid.UUID) (m []models.OperatingActivity, err error) {
	if m, err = module.db.operatingActivityModel.GetAllOperatingActivity(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllOperatingActivityPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.operatingActivityModel.GetAllOperatingActivityPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveOperatingActivity(id uuid.UUID) (m models.OperatingActivity, err error) {
	if m, err = module.db.operatingActivityModel.GetOneOperatingActivityByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertOperatingActivity(p *dto.InsertOperatingActivity) (err error) {
	if err = module.db.operatingActivityModel.InsertOperatingActivity(models.OperatingActivity{
		TaxInvoiceNumber:           p.TaxInvoiceNumber,
		DeliveryReceiptNumber:      p.DeliveryReceiptNumber,
		ClientID:                   p.ClientID,
		OperatingActivityCreatedBy: p.CreatedBy,
		Status:                     p.Status,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdateOperatingActivity(id uuid.UUID, p *dto.UpdateOperatingActivity) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	if p.Status == "done" {
		searchProductHistoryByOprActID := models.ProductHistory{
			OperatingActivityID: id,
		}

		var productHistory []models.ProductHistory
		err := tx.Model(&productHistory).Where(searchProductHistoryByOprActID).Find(&productHistory)

		if err.Error != nil {
			return err.Error
		}

		if len(productHistory) > 0 {
			for productHistoryIdx := 0; productHistoryIdx < len(productHistory); productHistoryIdx++ {
				productHistorySingle := productHistory[productHistoryIdx]

				// initial stock of the product
				initialProduct := models.Product{}

				if err := tx.Model(&initialProduct).First(&initialProduct, productHistorySingle.ProductID); err.Error != nil {
					return err.Error
				}

				productHistoryOutCount := 0.0
				productHistoryInCount := initialProduct.Stock
				productHistoryCountAccumulated := 0.0

				if productHistorySingle.Status == "increase" {
					productHistoryInCount += productHistorySingle.Quantity
				}

				if productHistorySingle.Status == "decrease" {
					productHistoryOutCount += productHistorySingle.Quantity
				}

				if (productHistoryInCount - productHistoryOutCount) <= 0 {
					productHistoryCountAccumulated += 0.0
				} else {
					productHistoryCountAccumulated = productHistoryInCount - productHistoryOutCount
					productUpdateQuery := models.Product{
						ID:    productHistorySingle.ProductID,
						Stock: productHistoryCountAccumulated,
					}

					if err := tx.Updates(&productUpdateQuery); err.Error != nil {
						tx.Rollback()
						return err.Error
					}
				}
			}
		}
	}

	operatingActivityUpdateQuery := models.OperatingActivity{
		ID:                         id,
		TaxInvoiceNumber:           p.TaxInvoiceNumber,
		DeliveryReceiptNumber:      p.DeliveryReceiptNumber,
		ClientID:                   p.ClientID,
		OperatingActivityUpdatedBy: p.UpdatedBy,
		Status:                     p.Status,
	}

	if err := tx.Updates(&operatingActivityUpdateQuery); err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	tx.Commit()
	return
}

func (module *module) CheckExistingOperatingActivity(id string, param CheckExistingOperatingActivityStruct) (err error) {

	if param.TaxInvoiceNumber != "" {
		if _, dbErr := module.db.operatingActivityModel.GetOneOperatingActivityByTaxNumber(param.TaxInvoiceNumber); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.operatingActivityModel.GetOneOperatingActivityByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
