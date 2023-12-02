package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/config"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(p dto.UserLogin) (token string, err error)

	CheckExistingUser(id string, param CheckExistingUserStruct) (err error)
	RegisterUser(p *dto.UserSignup) (err error)
	RegisterUserSuperAdmin(p dto.UserSignupSuperAdmin) (err error)
	RetrieveUser(id uuid.UUID) (m masterModels.User, err error)
	UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error)

	CheckExistingCompany(id string, param CheckExistingCompanyStruct) (err error)
	RetrieveAllCompany(id uuid.UUID) (m []masterModels.Company, err error)
	RetrieveCompany(id uuid.UUID) (m masterModels.Company, err error)
	InsertCompany(p *dto.InsertCompany) (err error)
	UpdateCompany(id uuid.UUID, p *dto.UpdateCompany) (err error)

	CheckExistingClient(id string, param CheckExistingClientStruct) (err error)
	RetrieveAllClient(id uuid.UUID) (m []masterModels.Client, err error)
	RetrieveClient(id uuid.UUID) (m masterModels.Client, err error)
	InsertClient(p *dto.InsertClient) (err error)
	UpdateClient(id uuid.UUID, p *dto.UpdateClient) (err error)

	CheckExistingOperatingActivity(id string, param CheckExistingOperatingActivityStruct) (err error)
	RetrieveAllOperatingActivity(id uuid.UUID) (m []models.OperatingActivity, err error)
	RetrieveOperatingActivity(id uuid.UUID) (m models.OperatingActivity, err error)
	InsertOperatingActivity(p *dto.InsertOperatingActivity) (err error)
	UpdateOperatingActivity(id uuid.UUID, p *dto.UpdateOperatingActivity) (err error)

	CheckExistingProduct(id string, param CheckExistingProductStruct) (err error)
	RetrieveAllProduct(id uuid.UUID) (m []models.Product, err error)
	RetrieveProduct(id uuid.UUID) (m models.Product, err error)
	InsertProduct(p *dto.InsertProduct) (err error)
	UpdateProduct(id uuid.UUID, p *dto.UpdateProduct) (err error)

	CheckExistingQuotation(id string, param CheckExistingQuotationStruct) (err error)
	RetrieveAllQuotation(id uuid.UUID) (m []models.Quotation, err error)
	RetrieveQuotation(id uuid.UUID) (m models.Quotation, err error)
	InsertQuotation(p *dto.InsertQuotation) (err error)
	UpdateQuotation(id uuid.UUID, p *dto.UpdateQuotation) (err error)

	CheckExistingProductHistory(id string, param CheckExistingProductHistoryStruct) (err error)
	RetrieveAllProductHistory(id uuid.UUID) (m []models.ProductHistory, err error)
	RetrieveProductHistory(id uuid.UUID) (m models.ProductHistory, err error)
	InsertProductHistory(p *dto.InsertProductHistory) (err error)
	UpdateProductHistory(id uuid.UUID, p *dto.UpdateProductHistory) (err error)

	CheckExistingPayment(id string, param CheckExistingPaymentStruct) (err error)
	RetrieveAllPayment(id uuid.UUID) (m []models.Payment, err error)
	RetrievePayment(id uuid.UUID) (m models.Payment, err error)
	InsertPayment(p *dto.InsertPayment) (err error)
	UpdatePayment(id uuid.UUID, p *dto.UpdatePayment) (err error)

	CheckExistingPaymentInstallment(id string, param CheckExistingPaymentInstallmentStruct) (err error)
	RetrieveAllPaymentInstallment(id uuid.UUID) (m []models.PaymentInstallment, err error)
	RetrievePaymentInstallment(id uuid.UUID) (m models.PaymentInstallment, err error)
	InsertPaymentInstallment(p *dto.InsertPaymentInstallment) (err error)
	UpdatePaymentInstallment(id uuid.UUID, p *dto.UpdatePaymentInstallment) (err error)

	CheckExistingOperatingActivityProduct(id string, param CheckExistingOperatingActivityProductStruct) (err error)
	RetrieveAllOperatingActivityProduct(id uuid.UUID) (m []models.OperatingActivityProduct, err error)
	RetrieveManyOperatingActivityProductByOperatingActivityID(id uuid.UUID) (m []models.OperatingActivityProduct, err error)
	RetrieveOperatingActivityProduct(id uuid.UUID) (m models.OperatingActivityProduct, err error)
	InsertOperatingActivityProduct(p *dto.InsertOperatingActivityProduct) (err error)
	UpdateOperatingActivityProduct(id uuid.UUID, p *dto.UpdateOperatingActivityProduct) (err error)

	CheckExistingPurchaseOrder(id string, param CheckExistingPurchaseOrderStruct) (err error)
	RetrieveAllPurchaseOrder(id uuid.UUID) (m []models.PurchaseOrder, err error)
	RetrievePurchaseOrder(id uuid.UUID) (m models.CustomResponsePurchaseOrder, err error)
	InsertPurchaseOrder(c *gin.Context, p *dto.InsertPurchaseOrder) (err error)
	UpdatePurchaseOrder(c *gin.Context, id uuid.UUID, p *dto.UpdatePurchaseOrder) (err error)

	CheckExistingPurchaseOrderProduct(id string, param CheckExistingPurchaseOrderProductStruct) (err error)
	RetrieveAllPurchaseOrderProduct(id uuid.UUID) (m []models.PurchaseOrderProduct, err error)
	RetrievePurchaseOrderProduct(id uuid.UUID) (m models.PurchaseOrderProduct, err error)
	InsertPurchaseOrderProduct(p *dto.InsertPurchaseOrderProduct) (err error)
	UpdatePurchaseOrderProduct(id uuid.UUID, p *dto.UpdatePurchaseOrderProduct) (err error)

	CheckExistingDeliveryOrder(id string, param CheckExistingDeliveryOrderStruct) (err error)
	RetrieveDeliveryOrder(id uuid.UUID) (m models.DeliveryOrder, err error)
	RetrieveAllDeliveryOrder(id uuid.UUID) (m []models.DeliveryOrder, err error)
	InsertDeliveryOrder(c *gin.Context, p *dto.InsertDeliveryOrder) (err error)
	UpdateDeliveryOrder(c *gin.Context, id uuid.UUID, p *dto.UpdateDeliveryOrder) (err error)

	CheckExistingInvoice(id string, param CheckExistingInvoiceStruct) (err error)
	RetrieveInvoice(id uuid.UUID) (m models.Invoice, err error)
	RetrieveAllInvoice(id uuid.UUID) (m []models.Invoice, err error)
	InsertInvoice(p *dto.InsertInvoice) (err error)
	UpdateInvoice(id uuid.UUID, p *dto.UpdateInvoice) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn                          *gorm.DB
	userModel                     masterModels.UserModelAction
	companyModel                  masterModels.CompanyModelAction
	clientModel                   masterModels.ClientModelAction
	operatingActivityModel        models.OperatingActivityModelAction
	productModel                  models.ProductModelAction
	productHistoryModel           models.ProductHistoryModelAction
	quotationModel                models.QuotationModelAction
	paymentModel                  models.PaymentModelAction
	paymentInstallmentModel       models.PaymentInstallmentModelAction
	operatingActivityProductModel models.OperatingActivityProductModelAction
	purchaseOrderModel            models.PurchaseOrderModelAction
	documentModel                 models.DocumentModelAction
	purchaseOrderProductModel     models.PurchaseOrderProductModelAction
	deliveryOrderModel            models.DeliveryOrderModelAction
	invoiceModel                  models.InvoiceModelAction
}

func InitializeServices() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(config.AppConfig.DBUrl), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return
	}
	log.Println("[INIT] connected to PostgresSQL")

	Handler = &module{
		db: &dbEntity{
			conn:                          db,
			userModel:                     masterModels.NewUserAction(db),
			companyModel:                  masterModels.NewCompanyAction(db),
			clientModel:                   masterModels.NewClientAction(db),
			operatingActivityModel:        models.NewOperatingActivityAction(db),
			productModel:                  models.NewProductAction(db),
			quotationModel:                models.NewQuotationAction(db),
			productHistoryModel:           models.NewProductHistoryAction(db),
			paymentModel:                  models.NewPaymentAction(db),
			paymentInstallmentModel:       models.NewPaymentInstallmentAction(db),
			operatingActivityProductModel: models.NewOperatingActivityProductAction(db),
			purchaseOrderModel:            models.NewPurchaseOrderAction(db),
			documentModel:                 models.NewDocumentAction(db),
			purchaseOrderProductModel:     models.NewPurchaseOrderProductAction(db),
			deliveryOrderModel:            models.NewDeliveryOrderAction(db),
			invoiceModel:                  models.NewInvoiceAction(db),
		},
	}
	return
}
