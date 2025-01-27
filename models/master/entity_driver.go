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

type DriverOrm struct {
	db *gorm.DB
}

type Driver struct {
	ID            uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name          string    `json:"name" gorm:"not null" binding:"required"`
	Phone         string    `json:"phone" gorm:"default:NULL" binding:"required"`
	PlateNumber   string    `json:"plateNumber,omitempty" gorm:"default:NULL"`
	LicensePhoto  string    `json:"licensePhoto,omitempty" gorm:"default:NULL"`
	DriverType    string    `json:"driverType,omitempty" gorm:"default:internal"`                              // external | internal | internal-agent
	Status        string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`       // active, inactive
	BookingStatus string    `json:"bookingStatus,omitempty" binding:"required" gorm:"not null;default:ready;"` // ready, busy
	Prob          int       `json:"prob,omitempty" binding:"required" gorm:"not null;default:0;"`

	CredentialsID uuid.UUID `json:"credentialsId" gorm:"type:uuid;not null"`
	CompanyID     uuid.UUID `json:"companyId" gorm:"type:uuid;default:NULL"`
	CreatedBy     uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy     uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Credentials   *Credentials     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CredentialsID;references:ID" json:"credentials,omitempty"`
	Company       *Company         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID;references:ID" json:"company,omitempty"`
	CarManagement []*CarManagement `json:"carManagement,omitempty"`

	types.DefaultModelProperty
}

type DriverModelAction interface {
	GetOneByID(id uuid.UUID) (m Driver, err error)
	GetOneByDriverName(Drivername string) (m Driver, err error)
	GetOneByEmail(email string) (m Driver, err error)
	GetOneMainAgentByCompanyId(companyId uuid.UUID) (Driver Driver, err error)
	GetAllDriverPaginated(c *gin.Context, DriverId uuid.UUID) (*database.Pagination, error)
	InsertDriver(p Driver, tx *gorm.DB) (err error)
	GetAllAvailableDriversByCompanyId(companyId uuid.UUID) (m []*Driver, err error)
	GetAllAvailableCompanyManagerByCompanyId() (m []*Driver, err error)
	GetOneByDriverID(id uuid.UUID) (Driver Driver, err error)

	UpdateDriverToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	UpdateDriverByCompanyId(companyId uuid.UUID, p Driver, tx *gorm.DB) (err error)
	UpdateDriverToBusy(id uuid.UUID, tx *gorm.DB) (err error)
	UpdateDriverToReady(id uuid.UUID, tx *gorm.DB) (err error)
	RemoveDriverFromCompany(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteDriver(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewDriverAction(db *gorm.DB) DriverModelAction {
	return &DriverOrm{db}
}

func (o *DriverOrm) GetAllDriverPaginated(c *gin.Context, DriverId uuid.UUID) (*database.Pagination, error) {
	var mArr []*Driver
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", DriverId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *DriverOrm) GetOneByID(id uuid.UUID) (Driver Driver, err error) {
	result := o.db.Model(&Driver).
		Where("credentials_id = ?", id).
		Preload(clause.Associations).
		First(&Driver)
	return Driver, result.Error
}

func (o *DriverOrm) GetOneByDriverID(id uuid.UUID) (Driver Driver, err error) {
	result := o.db.Model(&Driver).
		Where("id = ?", id).
		Preload("Company").
		First(&Driver)
	return Driver, result.Error
}

func (o *DriverOrm) GetOneMainAgentByCompanyId(companyId uuid.UUID) (Driver Driver, err error) {
	result := o.db.Model(&Driver).
		Where("company_id = ? AND driver_type = ?", companyId, "internal-agent").
		Preload(clause.Associations).
		First(&Driver)
	return Driver, result.Error
}

func (o *DriverOrm) GetAllAvailable(id uuid.UUID) (Driver []*Driver, err error) {
	result := o.db.Model(&Driver).
		Where("status = ? AND booking_status", "active", "ready").
		Preload(clause.Associations).
		Find(&Driver)
	return Driver, result.Error
}

func (o *DriverOrm) GetAllAvailableDriversByCompanyId(companyId uuid.UUID) (m []*Driver, err error) {
	result := o.db.Model(&m).Where("company_id = ? AND status = ?", companyId, "active").Find(&m)
	return m, result.Error
}

func (o *DriverOrm) GetAllAvailableCompanyManagerByCompanyId() (m []*Driver, err error) {
	result := o.db.Model(&m).Where("status = ? AND driver_type = ?", "active", "internal-agent").Find(&m)
	return m, result.Error
}

func (o *DriverOrm) GetOneByEmail(email string) (m Driver, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *DriverOrm) GetOneByDriverName(name string) (m Driver, err error) {
	result := o.db.Model(&m).Where("name = ?", name).First(&m)

	if m.Status == "inactive" {
		return m, errors.New("Driver status is inactive. please ask your administartor for further information")
	}
	return m, result.Error
}

func (o *DriverOrm) UpdateDriverToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Where("id = ?", id).Update("status", "inactive")
	return result.Error
}

func (o *DriverOrm) UpdateDriverByCompanyId(companyId uuid.UUID, p Driver, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Where("company_id = ?", companyId).Updates(&p)
	return result.Error
}

func (o *DriverOrm) UpdateDriverToBusy(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Where("id = ?", id).Update("booking_status", "busy")
	return result.Error
}

func (o *DriverOrm) UpdateDriverToReady(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Where("id = ?", id).Update("booking_status", "ready")
	return result.Error
}

func (o *DriverOrm) RemoveDriverFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Where("id", id).Update("company_id", nil)
	return result.Error
}

func (o *DriverOrm) InsertDriver(p Driver, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *DriverOrm) DeleteDriver(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Driver{}).Delete(&Driver{}, id)
	return result.Error
}
