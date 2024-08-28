package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BalanceDriverOrm struct {
	db *gorm.DB
}

type BalanceDriver struct {
	ID     uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Amount float64   `gorm:"not null;default:0" json:"amount,omitempty" binding:"required"`

	DriverID  uuid.UUID `json:"driverId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver *master.Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`

	types.DefaultModelProperty
}

type BalanceDriverModelAction interface {
	GetOneByID(id uuid.UUID) (m BalanceDriver, err error)
	GetOneByEmail(email string) (m BalanceDriver, err error)

	InsertBalanceDriver(p BalanceDriver) (err error)
	DeleteBalanceDriver(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBalanceDriverAction(db *gorm.DB) BalanceDriverModelAction {
	return &BalanceDriverOrm{db}
}

func (o *BalanceDriverOrm) GetOneByID(id uuid.UUID) (BalanceDriver BalanceDriver, err error) {
	result := o.db.Model(&BalanceDriver).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&BalanceDriver)
	return BalanceDriver, result.Error
}

func (o *BalanceDriverOrm) GetOneByEmail(email string) (m BalanceDriver, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BalanceDriverOrm) InsertBalanceDriver(p BalanceDriver) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *BalanceDriverOrm) DeleteBalanceDriver(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BalanceDriver{}).Delete(&BalanceDriver{}, id)
	return result.Error
}
