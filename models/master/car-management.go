package master

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarManagementOrm struct {
	db *gorm.DB
}

type CarManagement struct {
	ID                uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name              string    `json:"name" gorm:"not null" binding:"required"`
	PlateNumber       string    `json:"plateNumber,omitempty" gorm:"default:NULL"`
	FrontCarPhoto     string    `json:"frontCarPhoto,omitempty" gorm:"default:NULL"`
	LicensePhoto      string    `json:"licensePhoto,omitempty" gorm:"default:NULL"`
	RoadTaxPhoto      string    `json:"roadTaxPhoto,omitempty" gorm:"default:NULL"`
	VEPPhoto          string    `json:"vepPhoto,omitempty" gorm:"default:NULL"`
	CarManagementType string    `json:"carType,omitempty" gorm:"default:external"`
	FileName          string    `json:"fileName,omitempty" gorm:"default:NULL"`
	Status            string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	CarModelID uuid.UUID `json:"carModelId" gorm:"type:uuid;not null"`
	DriverID   uuid.UUID `json:"driverId" gorm:"type:uuid;not null"`
	CompanyID  uuid.UUID `json:"companyId" gorm:"type:uuid;default:NULL"`
	CreatedBy  uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy  uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver   *Driver   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	CarModel *CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarModelID;references:ID" json:"carModel,omitempty"`
	Company  *Company  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID;references:ID" json:"company,omitempty"`

	types.DefaultModelProperty
}

type CarManagementModelAction interface {
	GetOneByID(id uuid.UUID) (m CarManagement, err error)
	GetOneByCarManagementName(CarManagementname string) (m CarManagement, err error)
	GetOneByEmail(email string) (m CarManagement, err error)
	GetAllCarManagementPaginated(c *gin.Context, CarManagementId uuid.UUID) (*database.Pagination, error)
	GetOneByCarModelIDAndAvailable(id uuid.UUID, driverId uuid.UUID, companyId uuid.UUID) (CarManagement CarManagement, err error)
	GetAllCarManagementByDriverID(driverId uuid.UUID) (CarManagement []CarManagement, err error)
	UpdateInternalCarManagementStatus(plateNumber string, status string, tx *gorm.DB) (err error)
	GetOneByDriverIdAndCompanyId(driverId uuid.UUID, companyId uuid.UUID) (m CarManagement, err error)
	GetAllAvailableCarManagementByCompanyIdAndDriverId(companyId uuid.UUID, driverId uuid.UUID) (CarManagement []*CarManagement, err error)
	GetSameCarManagementByAgentByPlateNumberAndCompanyIdAndDriverId(companyId uuid.UUID, plateNumber string, driverId uuid.UUID) (CarManagement CarManagement, err error)
	CheckKhaimalDriverAvailable(carModelId uuid.UUID) (CarManagement []*CarManagement, err error)
	GetKhaimalManagerId(carModelId uuid.UUID) (CarManagement CarManagement, err error)

	InsertCarManagement(p CarManagement) (err error)
	UpdateCarManagementByCompanyId(companyId uuid.UUID, p CarManagement, tx *gorm.DB) (err error)
	UpdateCarManagementToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	RemoveCarManagementFromCompany(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteCarManagement(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewCarManagementAction(db *gorm.DB) CarManagementModelAction {
	return &CarManagementOrm{db}
}

func (o *CarManagementOrm) GetAllCarManagementPaginated(c *gin.Context, CarManagementId uuid.UUID) (*database.Pagination, error) {
	var mArr []*CarManagement
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", CarManagementId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *CarManagementOrm) GetOneByID(id uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) CheckKhaimalDriverAvailable(carModelId uuid.UUID) (CarManagement []*CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		// Where("car_model_id = ? AND status = ? AND company_id = ?", carModelId, "active", "45d64f59-bdc4-48cd-853b-33d88182e065").
		Where("car_model_id = ? AND status = ? AND company_id = ?", carModelId, "active", "1af117b6-dc78-4fb1-abab-2a9645651a88").
		// Preload(clause.Associations).
		Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credentials")
		}).
		Preload("Company").
		Preload("CarModel").
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetKhaimalManagerId(carModelId uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		// Where("car_model_id = ? AND status = ? AND driver_id = ?", carModelId, "active", "45d64f59-bdc4-48cd-853b-33d88182e065").
		Where("car_model_id = ? AND status = ? AND driver_id = ?", carModelId, "active", "de329135-f7de-483f-8dbf-2a125d323da3").
		// Preload(clause.Associations).
		Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credentials")
		}).
		Preload("Company").
		Preload("CarModel").
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetAllAvailableCarManagementByCompanyIdAndDriverId(companyId uuid.UUID, driverId uuid.UUID) (CarManagement []*CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("driver_id = ? AND company_id = ?", driverId, companyId).
		Preload("Driver").
		Preload("CarModel", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_num DESC")
		}).
		Preload("Company").
		Find(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetSameCarManagementByAgentByPlateNumberAndCompanyIdAndDriverId(companyId uuid.UUID, plateNumber string, driverId uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("plate_number = ? AND company_id = ? AND driver_id = ?", plateNumber, companyId, driverId).
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetOneByDriverIdAndCarManagementId(id uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("driver_id = ?", id).
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetOneByCarModelIDAndAvailable(id uuid.UUID, driverId uuid.UUID, companyId uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("car_model_id = ? AND status = ? AND driver_id = ? AND company_id = ?", id, "active", driverId, companyId).
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetAllCarManagementByDriverID(driverId uuid.UUID) (CarManagement []CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("driver_id", driverId).
		Preload("CarModel").
		Find(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetOneByEmail(email string) (m CarManagement, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *CarManagementOrm) GetOneByDriverIdAndCompanyId(driverId uuid.UUID, companyId uuid.UUID) (m CarManagement, err error) {
	result := o.db.Model(&m).Where("driver_id = ? AND company_id = ?", driverId, companyId).First(&m)
	return m, result.Error
}

func (o *CarManagementOrm) GetOneByCarManagementName(name string) (m CarManagement, err error) {
	result := o.db.Model(&m).Where("name = ?", name).First(&m)
	if m.Status == "inactive" {
		return m, errors.New("CarManagement status is inactive. please ask your administartor for further information")
	}
	return m, result.Error
}

func (o *CarManagementOrm) UpdateCarManagementByCompanyId(companyId uuid.UUID, p CarManagement, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("company_id", companyId).Updates(&p)
	return result.Error
}

func (o *CarManagementOrm) UpdateCarManagementToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("id = ?", id).Update("status", "inactive")
	return result.Error
}

func (o *CarManagementOrm) RemoveCarManagementFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("id", id).Update("company_id", nil)
	return result.Error
}

func (o *CarManagementOrm) UpdateInternalCarManagementStatus(plateNumber string, status string, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("status = ? AND car_management_type = ? AND plate_number = ?", "active", "internal", plateNumber).Update("status", status)
	return result.Error
}

func (o *CarManagementOrm) InsertCarManagement(p CarManagement) (err error) {
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *CarManagementOrm) DeleteCarManagement(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Delete(&CarManagement{}, id)
	return result.Error
}
