package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	// AUTH
	AuthenticateUser(p dto.CredentialSignInDto) (token string, err error)
	AuthenticateUserV2(credentials dto.CredentialSignInDto) (token string, err error)
	CheckExistingUser(id string, param CheckExistingUserStruct) (err error)
	CheckExistingUserV2(id string, param CheckExistingUserStruct) (m *masterModels.Credentials, err error)
	RetrieveAllUserPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error)
	RetrieveUser(id uuid.UUID) (m masterModels.User, err error)
	DeleteUser(id uuid.UUID) (err error)
	UpdateDeviceToken(token string, credentialId uuid.UUID) (err error)
	ResetPassword(p *dto.UniversalResetPassword) (err error)
	// Company
	InsertCompany(c *gin.Context, p *dto.InsertCompany, pAgent *dto.DriverSignUpValidator, pDrivers []*dto.DriverSignUpValidator, pTransports []*dto.InsertCarManagement) (err error)
	ApproveCompany(companyId uuid.UUID) (err error)
	RetrieveAllAvailableDrivers(companyId uuid.UUID) (m []*masterModels.Driver, err error)
	RegisterNewDriverInternalAgent(driverId uuid.UUID, credentials *dto.DriverSignup) (err error)
	RetrieveAllRegisteredCustomerPaginated(c *gin.Context) (pagination *database.Pagination, err error)
	// Credentials
	RetrieveEntityCredentialsByUserID(userId uuid.UUID) (m masterModels.Credentials, err error)
	RetrieveEntityCredentialsByEmail(email string) (m master.Credentials, err error)
	// Booking Transfer
	RetrieveLastOrderByCustomerID(userId uuid.UUID) (m models.BookingTransfer, err error)
	RetrieveAllKhaimalBookingTransfer() (m []*models.BookingTransfer, err error)
	InsertBookingTransfer(p *dto.InsertBookingTransfer) (err error)
	UpdateBookingTransfer(id uuid.UUID, p *dto.UpdateBookingTransfer) (err error)
	RetrieveAllBookingTransferByCustomer(id uuid.UUID) (m []*models.BookingTransfer, err error)
	CustomerCancelBooking(id uuid.UUID) (err error)
	// Booking Transfer - Assigned
	RetrieveBookingTransferAssignedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error)
	RetrieveBookingTransferAcceptedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error)
	RetrieveBookingTransferCancelledByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error)
	RetrieveBookingTransferOngoingByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error)
	RetrieveBookingTransferCompletedByDriverID(userId uuid.UUID) (m []models.BookingTransferAssigned, err error)
	AcceptBookingTransfer(id uuid.UUID) (err error)
	CancelBookingTransfer(id uuid.UUID) (err error)
	OngoingBookingTransfer(id uuid.UUID) (err error)
	CompleteBookingTransfer(id uuid.UUID) (err error)
	CompletePickupBooking(id uuid.UUID) (err error)
	SwitchDriver(id uuid.UUID, driverId uuid.UUID, plateNumber string) (err error)
	// Booking Transfer - Rating
	InsertBookingTransferRating(p *dto.InsertBookingTransferRating) (err error)
	// CAR MANAGEMENT
	InsertCarManagement(c *gin.Context, p *dto.InsertCarManagement) (err error)
	UpdateCarManagement(c *gin.Context, p *dto.UpdateCarManagement, id uuid.UUID) (err error)
	DeleteCarManagement(c *gin.Context, id uuid.UUID) (err error)
	RetrieveCarManagementrByUserID(userId uuid.UUID) (m []masterModels.CarManagement, err error)
	SetCarManagementInactive(c *gin.Context, id uuid.UUID) (err error)
	SetCarManagementActive(c *gin.Context, id uuid.UUID) (err error)
	RetrieveAllAvailable(userId uuid.UUID) (m []masterModels.CarManagement, err error)
	// AUTH - CUSTOMER
	UpdateCustomerByCredID(c *gin.Context, p *dto.UpdateCustomer, id uuid.UUID) (err error)
	RegisterCustomer(credentials *dto.CustomerSignup) (err error)
	RetrieveEntityCustomerByUserID(userId uuid.UUID) (m masterModels.Customer, err error)
	// AUTH - INTERNAL
	RetrieveEntityInternalByUserID(userId uuid.UUID) (m masterModels.Internal, err error)
	// AUTH - DRIVER
	RegisterDriver(credentials *dto.DriverSignup) (err error)
	RetrieveDriverLinkedCompany(driverId uuid.UUID) (m master.Driver, err error)
	// AUTH - PUBLIC
	RegisterUser(p *dto.UserSignup) (err error)
	// WEB STATISTIC - CUSTOMER
	RetrieveCustomerWebStatisticByUserID(userId uuid.UUID) (ctx int64, err error)
	RetrieveCustomerWebStatisticActiveByUserID(userId uuid.UUID) (ctx int64, err error)
	// DRIVER TOPUP
	RetrieveAllDriverTopup() (m []models.DriverTopup, err error)
	InsertDriverTopup(c *gin.Context, p *dto.InsertDriverTopup) (err error)
	// DRIVER TOPUP HISTORY
	RetrieveAllDriverTopupByDriver(id uuid.UUID) (m []models.DriverTransactionHistory, err error)
	ApproveDriverTopup(c *gin.Context, p *dto.UpdateDriverTopupHistory) (err error)
	RejectDriverTopup(c *gin.Context, p *dto.UpdateDriverTopupHistory) (err error)
	// DRIVER BALANCE
	RetrieveDriverBalanceDetailByDriver(id uuid.UUID) (m models.BalanceDriver, err error)
	// PAYMENT
	InsertPayment(p *models.Payment) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn                          *gorm.DB
	userModel                     masterModels.UserModelAction
	paymentModel                  models.PaymentModelAction
	credentialModel               masterModels.CredentialsModelAction
	driverModel                   masterModels.DriverModelAction
	balanceDriver                 models.BalanceDriverModelAction
	companyModel                  masterModels.CompanyModelAction
	userCustomerModel             masterModels.CustomerModelAction
	userDriverModel               masterModels.DriverModelAction
	userInternalModel             masterModels.InternalModelAction
	carManagementModel            masterModels.CarManagementModelAction
	carModel                      masterModels.CarModelModelAction
	bookingTransfer               models.BookingTransferModelAction
	bookingTransferAssigned       models.BookingTransferAssignedModelAction
	bookingTransferRating         models.BookingTransferRatingModelAction
	driverTopupModel              models.DriverTopupModelAction
	driverTransactionHistoryModel models.DriverTransactionHistoryModelAction
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
			conn:                          db,
			paymentModel:                  models.NewPaymentAction(db),
			userModel:                     masterModels.NewUserAction(db),
			credentialModel:               masterModels.NewCredentialsAction(db),
			driverModel:                   masterModels.NewDriverAction(db),
			companyModel:                  masterModels.NewCompanyAction(db),
			userCustomerModel:             masterModels.NewCustomerAction(db),
			userDriverModel:               masterModels.NewDriverAction(db),
			userInternalModel:             masterModels.NewInternalAction(db),
			carManagementModel:            masterModels.NewCarManagementAction(db),
			carModel:                      masterModels.NewCarModelAction(db),
			bookingTransfer:               models.NewBookingTransferAction(db),
			balanceDriver:                 models.NewBalanceDriverAction(db),
			bookingTransferAssigned:       models.NewBookingTransferAssignedAction(db),
			bookingTransferRating:         models.NewBookingTransferRatingAction(db),
			driverTopupModel:              models.NewDriverTopupAction(db),
			driverTransactionHistoryModel: models.NewDriverTransactionHistoryAction(db),
		},
	}
	return
}
