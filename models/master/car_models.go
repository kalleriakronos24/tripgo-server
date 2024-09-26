package master

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarModelOrm struct {
	db *gorm.DB
}

type CarModel struct {
	ID            uuid.UUID        `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name          string           `json:"name" gorm:"not null" binding:"required"`
	LuggageCount  int              `json:"luggageCount,omitempty" gorm:"not null" binding:"required"`
	PersonCount   int              `json:"personCount,omitempty" gorm:"not null" binding:"required"`
	ImagePath     string           `json:"imagePath,omitempty" gorm:"not null" binding:"required"`
	Status        string           `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`
	CreatedBy     uuid.UUID        `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy     uuid.UUID        `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`
	CarManagement []*CarManagement `json:"carManagement,omitempty"`
	types.DefaultModelProperty
}

type CarModelModelAction interface {
	GetOneByID(id uuid.UUID) (m CarModel, err error)
	GetAllCarModelPaginated(c *gin.Context, CarModelId uuid.UUID) (*database.Pagination, error)
	InsertCarModel(p CarModel) (err error)
	UpdateCarModelToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteCarModel(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewCarModelAction(db *gorm.DB) CarModelModelAction {
	return &CarModelOrm{db}
}

func (o *CarModelOrm) GetAllCarModelPaginated(c *gin.Context, CarModelId uuid.UUID) (*database.Pagination, error) {
	var mArr []*CarModel
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", CarModelId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
}

func (o *CarModelOrm) GetOneByID(id uuid.UUID) (CarModel CarModel, err error) {
	result := o.db.Model(&CarModel).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&CarModel)
	return CarModel, result.Error
}

func (o *CarModelOrm) UpdateCarModelToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarModel{}).Where("id = ?", id).Update("status", "inactive")
	return result.Error
}

func (o *CarModelOrm) InsertCarModel(p CarModel) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *CarModelOrm) DeleteCarModel(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&CarModel{}).Delete(&CarModel{}, id)
	return result.Error
}
