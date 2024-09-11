package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	// AUTH
	AuthenticateUser(p dto.CredentialSignInDto) (token string, err error)
	CheckExistingUser(id string, param CheckExistingUserStruct) (err error)
	RetrieveAllUserPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error)
	RetrieveUser(id uuid.UUID) (m masterModels.User, err error)
	DeleteUser(id uuid.UUID) (err error)
	// AUTH - CUSTOMER
	RegisterCustomer(credentials *dto.CustomerSignup) (err error)
	// AUTH - DRIVER
	RegisterDriver(credentials *dto.DriverSignup) (err error)
	// AUTH - PUBLIC
	RegisterUser(p *dto.UserSignup) (err error)
	// OTHER
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn               *gorm.DB
	userModel          masterModels.UserModelAction
	credentialModel    masterModels.CredentialsModelAction
	companyModel       masterModels.CompanyModelAction
	userCustomerModel  masterModels.CustomerModelAction
	userDriverModel    masterModels.DriverModelAction
	userInternalModel  masterModels.InternalModelAction
	carManagementModel masterModels.CarManagementModelAction
	carModel           masterModels.CarModelModelAction
}

type GenerateDocumentOutput struct {
	OutputPath string
	FileName   string
}

func InitializeServices() (err error) {
	// Initialize DB
	db := database.GetDatabaseConnection()

	Handler = &module{
		db: &dbEntity{
			conn:               db,
			userModel:          masterModels.NewUserAction(db),
			credentialModel:    masterModels.NewCredentialsAction(db),
			companyModel:       masterModels.NewCompanyAction(db),
			userCustomerModel:  masterModels.NewCustomerAction(db),
			userDriverModel:    masterModels.NewDriverAction(db),
			userInternalModel:  masterModels.NewInternalAction(db),
			carManagementModel: masterModels.NewCarManagementAction(db),
			carModel:           masterModels.NewCarModelAction(db),
		},
	}
	return
}
