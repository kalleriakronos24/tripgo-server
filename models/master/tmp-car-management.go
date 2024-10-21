package master

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TempCarManagementOrm struct {
	db *gorm.DB
}

type TempCarManagement struct {
	ID                    uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name                  string    `json:"name" gorm:"not null" binding:"required"`
	FrontCarPhoto         string    `json:"frontCarPhoto,omitempty" gorm:"default:NULL"`
	Luggage               int       `json:"luggage,omitempty" gorm:"default:NULL"`
	LicensePhoto          string    `json:"licensePhoto,omitempty" gorm:"default:NULL"`
	TempCarManagementType string    `json:"TempCarManagementType,omitempty" gorm:"default:external"`
	Status                string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	DriverID  uuid.UUID `json:"driverId" gorm:"type:uuid;not null"`
	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver *Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	types.DefaultModelProperty
}

type TempCarManagementModelAction interface {
	GetOneByID(id uuid.UUID) (m TempCarManagement, err error)
	GetOneByTempCarManagementName(TempCarManagementname string) (m TempCarManagement, err error)
	GetOneByEmail(email string) (m TempCarManagement, err error)
	GetAllTempCarManagementPaginated(c *gin.Context, TempCarManagementId uuid.UUID) (*database.Pagination, error)

	InsertTempCarManagement(p TempCarManagement) (err error)

	UpdateTempCarManagementToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	RemoveTempCarManagementFromCompany(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteTempCarManagement(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewTempCarManagementAction(db *gorm.DB) TempCarManagementModelAction {
	return &TempCarManagementOrm{db}
}

func (o *TempCarManagementOrm) GetAllTempCarManagementPaginated(c *gin.Context, TempCarManagementId uuid.UUID) (*database.Pagination, error) {
	var mArr []*TempCarManagement
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", TempCarManagementId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *TempCarManagementOrm) GetOneByID(id uuid.UUID) (TempCarManagement TempCarManagement, err error) {
	result := o.db.Model(&TempCarManagement).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&TempCarManagement)
	return TempCarManagement, result.Error
}

func (o *TempCarManagementOrm) GetOneByEmail(email string) (m TempCarManagement, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *TempCarManagementOrm) GetOneByTempCarManagementName(name string) (m TempCarManagement, err error) {
	result := o.db.Model(&m).Where("name = ?", name).First(&m)

	if m.Status == "inactive" {
		return m, errors.New("TempCarManagement status is inactive. please ask your administartor for further information")
	}

	return m, result.Error
}

func (o *TempCarManagementOrm) UpdateTempCarManagementToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&TempCarManagement{}).Where("id = ?", id).Update("status", "inactive")
	return result.Error
}

func (o *TempCarManagementOrm) RemoveTempCarManagementFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&TempCarManagement{}).Where("id", id).Update("company_id", nil)
	return result.Error
}

func (o *TempCarManagementOrm) InsertTempCarManagement(p TempCarManagement) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *TempCarManagementOrm) DeleteTempCarManagement(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&TempCarManagement{}).Delete(&TempCarManagement{}, id)
	return result.Error
}
