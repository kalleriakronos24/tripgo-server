package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gorm.io/gorm"
)

type CheckExistingPurchaseOrderProductStruct struct {
	*models.PurchaseOrderProduct
}

func (module *module) RetrieveAllPurchaseOrderProduct(id uuid.UUID) (m []models.PurchaseOrderProduct, err error) {
	if m, err = module.db.purchaseOrderProductModel.GetAllPurchaseOrderProduct(id); err != nil {
		return m, err
	}
	return
}

func (module *module) RetrieveAllPurchaseOrderProductPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.purchaseOrderProductModel.GetAllPurchaseOrderProductPaginated(c, id); err != nil {
		return pagination, err
	}
	return
}

func (module *module) RetrievePurchaseOrderProduct(id uuid.UUID) (m models.PurchaseOrderProduct, err error) {
	if m, err = module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByID(id); err != nil {
		return m, err
	}
	return
}

func (module *module) InsertPurchaseOrderProduct(p *dto.InsertPurchaseOrderProduct) (err error) {

	//tx := database.GetDatabaseConnection().Begin()
	//
	//purchaseOrder := models.PurchaseOrder{}
	//productModel := models.Product{}
	//
	//if err := tx.Model(&purchaseOrder).First(&purchaseOrder, p.PurchaseOrderID); err != nil {
	//	tx.Rollback()
	//	return err.Error
	//}
	//
	//if purchaseOrder.Type == "out" {
	//	if err := tx.Model(&productModel).First(p.ProductID); err != nil {
	//		tx.Rollback()
	//		return err.Error
	//	}
	//
	//	addedProductStock := productModel.Stock + p.Quantity
	//	updatePayload := models.Product{
	//		ID:    p.ProductID,
	//		Stock: addedProductStock,
	//	}
	//
	//	if err := tx.Model(&productModel).Updates(&updatePayload); err != nil {
	//		tx.Rollback()
	//		return err.Error
	//	}
	//}

	//purchaseOrderProductModel := models.PurchaseOrderProduct{
	//	Quantity:                      p.Quantity,
	//	VATRate:                       p.VATRate,
	//	SubTotal:                      p.SubTotal,
	//	GrandTotal:                    p.GrandTotal,
	//	PurchaseOrderID:               p.PurchaseOrderID,
	//	ClientID:                      p.ClientID,
	//	ProductID:                     p.ProductID,
	//	PurchaseOrderProductCreatedBy: p.CreatedBy,
	//}
	//
	//if err := tx.Model(&purchaseOrderProductModel).Create(&purchaseOrderProductModel); err != nil {
	//	tx.Rollback()
	//	return err.Error
	//}
	//tx.Commit()

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
		return err
	}
	return
}

func (module *module) UpdatePurchaseOrderProduct(id uuid.UUID, p *dto.UpdatePurchaseOrderProduct) (err error) {
	//tx := database.GetDatabaseConnection().Begin()

	//purchaseOrder := models.PurchaseOrder{}
	//productModel := models.Product{}
	//purchaseOrderProduct := models.PurchaseOrderProduct{}
	//
	//if err := tx.Model(&purchaseOrder).First(&purchaseOrder, p.PurchaseOrderID); err != nil {
	//	tx.Rollback()
	//	return err.Error
	//}
	//
	//if err := tx.Model(&purchaseOrderProduct).First(&purchaseOrderProduct, id); err != nil {
	//	tx.Rollback()
	//	return err.Error
	//}
	//
	//if purchaseOrder.Type == "out" {
	//
	//	if err := tx.Model(&productModel).First(p.ProductID); err != nil {
	//		tx.Rollback()
	//		return err.Error
	//	}
	//
	//	updatedProductStock := (productModel.Stock - purchaseOrderProduct.Quantity) + p.Quantity
	//	updatePayload := models.Product{
	//		ID:    p.ProductID,
	//		Stock: updatedProductStock,
	//	}
	//
	//	if err := tx.Model(&productModel).Updates(&updatePayload); err != nil {
	//		tx.Rollback()
	//		return err.Error
	//	}
	//}

	//purchaseOrderProductModel := models.PurchaseOrderProduct{
	//	ID:                            id,
	//	Quantity:                      p.Quantity,
	//	VATRate:                       p.VATRate,
	//	SubTotal:                      p.SubTotal,
	//	GrandTotal:                    p.GrandTotal,
	//	PurchaseOrderID:               p.PurchaseOrderID,
	//	ClientID:                      p.ClientID,
	//	ProductID:                     p.ProductID,
	//	PurchaseOrderProductUpdatedBy: p.UpdatedBy,
	//}
	//
	//if err := tx.Model(&purchaseOrderProductModel).Where(id).Updates(&purchaseOrderProductModel); err != nil {
	//	tx.Rollback()
	//	return err.Error
	//}
	//tx.Commit()

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
		return err
	}

	return
}

func (module *module) DeletePurchaseOrderProduct(id uuid.UUID, tx *gorm.DB) (err error) {
	if err = module.db.purchaseOrderProductModel.DeletePurchaseOrderProductByProductID(id, tx); err != nil {
		return err
	}
	return
}

func (module *module) CheckExistingPurchaseOrderProduct(id string, param CheckExistingPurchaseOrderProductStruct) (err error) {

	if param.PurchaseOrderID.String() != "" {
		if _, dbErr := module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByPurchaseOrderId(param.PurchaseOrderID); dbErr != nil {
			return dbErr
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.purchaseOrderProductModel.GetOnePurchaseOrderProductByID(uid); dbErr != nil {
			return dbErr
		}
	}
	return
}
