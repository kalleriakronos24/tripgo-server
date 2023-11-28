package services

import (
	"gitlab.com/odma1/odma-be/config"
	"gitlab.com/odma1/odma-be/models"
	"log"

	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn                   *gorm.DB
	userModel              masterModels.UserModelAction
	companyModel           masterModels.CompanyModelAction
	clientModel            masterModels.ClientModelAction
	operatingActivityModel models.OperatingActivityModelAction
	productModel           models.ProductModelAction
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
			conn:                   db,
			userModel:              masterModels.NewUserAction(db),
			companyModel:           masterModels.NewCompanyAction(db),
			clientModel:            masterModels.NewClientAction(db),
			operatingActivityModel: models.NewOperatingActivtyAction(db),
			productModel:           models.NewProductAction(db),
		},
	}
	return
}
