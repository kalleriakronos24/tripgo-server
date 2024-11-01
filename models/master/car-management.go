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
	CarManagementType string    `json:"carType,omitempty" gorm:"default:external"`
	FileName          string    `json:"fileName,omitempty" gorm:"default:NULL"`
	Status            string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	CarModelID uuid.UUID `json:"carModelId" gorm:"type:uuid;not null"`
	DriverID   uuid.UUID `json:"driverId" gorm:"type:uuid;not null"`
	CreatedBy  uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy  uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver   *Driver   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	CarModel *CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarModelID;references:ID" json:"carModel,omitempty"`

	types.DefaultModelProperty
}

type CarManagementModelAction interface {
	GetOneByID(id uuid.UUID) (m CarManagement, err error)
	GetOneByCarManagementName(CarManagementname string) (m CarManagement, err error)
	GetOneByEmail(email string) (m CarManagement, err error)
	GetAllCarManagementPaginated(c *gin.Context, CarManagementId uuid.UUID) (*database.Pagination, error)
	GetOneByCarModelIDAndAvailable(id uuid.UUID, driverId uuid.UUID) (CarManagement CarManagement, err error)
	GetAllCarManagementByDriverID(driverId uuid.UUID) (CarManagement []CarManagement, err error)

	InsertCarManagement(p CarManagement) (err error)
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

func (o *CarManagementOrm) GetOneByDriverIdAndCarManagementId(id uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("driver_id = ?", id).
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetOneByCarModelIDAndAvailable(id uuid.UUID, driverId uuid.UUID) (CarManagement CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("car_model_id = ? AND status = ?", id, "active").
		Preload(clause.Associations).
		First(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetAllCarManagementByDriverID(driverId uuid.UUID) (CarManagement []CarManagement, err error) {
	result := o.db.Model(&CarManagement).
		Where("driver_id", driverId).
		Preload(clause.Associations).
		Find(&CarManagement)
	return CarManagement, result.Error
}

func (o *CarManagementOrm) GetOneByEmail(email string) (m CarManagement, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *CarManagementOrm) GetOneByCarManagementName(name string) (m CarManagement, err error) {
	result := o.db.Model(&m).Where("name = ?", name).First(&m)

	if m.Status == "inactive" {
		return m, errors.New("CarManagement status is inactive. please ask your administartor for further information")
	}
	return m, result.Error
}

func (o *CarManagementOrm) UpdateCarManagementToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("id = ?", id).Update("status", "inactive")
	return result.Error
}

func (o *CarManagementOrm) RemoveCarManagementFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarManagement{}).Where("id", id).Update("company_id", nil)
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
