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

type CheckExistingPurchaseOrderProductStruct struct {
	*models.PurchaseOrderProduct
}

func (module *module) RetrieveAllPurchaseOrderProduct(id uuid.UUID) (m []models.PurchaseOrderProduct, err error) {
	if m, err = module.db.purchaseOrderProductModel.GetAllPurchaseOrderProduct(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllPurchaseOrderProductPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.purchaseOrderProductModel.GetAllPurchaseOrderProductPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrievePurchaseOrderProduct(id uuid.UUID) (m models.PurchaseOrderProduct, err error) {
	if m, err = module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertPurchaseOrderProduct(p *dto.InsertPurchaseOrderProduct) (err error) {
	if err = module.db.purchaseOrderProductModel.InsertPurchaseOrderProduct(models.PurchaseOrderProduct{
		Quantity:                      p.Quantity,
		VATRate:                       p.VATRate,
		SubTotal:                      p.SubTotal,
		GrandTotal:                    p.GrandTotal,
		PurchaseOrderID:               p.PurchaseOrderID,
		ClientID:                      p.ClientID,
		ProductID:                     p.ProductID,
		PurchaseOrderProductCreatedBy: p.CreatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) UpdatePurchaseOrderProduct(id uuid.UUID, p *dto.UpdatePurchaseOrderProduct) (err error) {
	if err = module.db.purchaseOrderProductModel.UpdatePurchaseOrderProduct(id, models.PurchaseOrderProduct{
		Quantity:                      p.Quantity,
		VATRate:                       p.VATRate,
		SubTotal:                      p.SubTotal,
		GrandTotal:                    p.GrandTotal,
		PurchaseOrderID:               p.PurchaseOrderID,
		ClientID:                      p.ClientID,
		ProductID:                     p.ProductID,
		PurchaseOrderProductUpdatedBy: p.UpdatedBy,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) CheckExistingPurchaseOrderProduct(id string, param CheckExistingPurchaseOrderProductStruct) (err error) {

	if param.PurchaseOrderID.String() != "" {
		if _, dbErr := module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByPurchaseOrderId(param.PurchaseOrderID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
