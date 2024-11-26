package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DriverTopupOrm struct {
	db *gorm.DB
}

type DriverTopup struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Uid          string    `json:"uid,omitempty" gorm:"default:NULL"`
	Amount       float64   `gorm:"not null;default:0" json:"amount,omitempty" binding:"required"`
	ReceiptImage string    `json:"receiptImage,omitempty" gorm:"default:NULL"`

	DriverID  uuid.UUID `json:"driverId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver                   *master.Driver            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	DriverTransactionHistory *DriverTransactionHistory `json:"transactionHistory,omitempty"`
	types.DefaultModelProperty
}

type DriverTopupModelAction interface {
	GetOneByID(id uuid.UUID) (m DriverTopup, err error)
	GetOneByEmail(email string) (m DriverTopup, err error)
	GetAllTopupRequest() (DriverTopup []DriverTopup, err error)

	InsertDriverTopup(p DriverTopup, tx *gorm.DB) (err error)
	DeleteDriverTopup(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewDriverTopupAction(db *gorm.DB) DriverTopupModelAction {
	return &DriverTopupOrm{db}
}

func (o *DriverTopupOrm) GetAllTopupRequest() (DriverTopup []DriverTopup, err error) {
	result := o.db.Model(&DriverTopup).
		Preload("Driver").
		Preload("DriverTransactionHistory").
		Find(&DriverTopup)
	return DriverTopup, result.Error
}

func (o *DriverTopupOrm) GetOneByID(id uuid.UUID) (DriverTopup DriverTopup, err error) {
	result := o.db.Model(&DriverTopup).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&DriverTopup)
	return DriverTopup, result.Error
}
func (o *DriverTopupOrm) GetOneByEmail(email string) (m DriverTopup, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *DriverTopupOrm) InsertDriverTopup(p DriverTopup, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *DriverTopupOrm) DeleteDriverTopup(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&DriverTopup{}).Delete(&DriverTopup{}, id)
	return result.Error
}
