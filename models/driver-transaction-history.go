package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DriverTransactionHistoryOrm struct {
	db *gorm.DB
}

type DriverTransactionHistory struct {
	ID     uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Remark string    `json:"remark,omitempty" gorm:"default:NULL"`
	Status string    `json:"status,omitempty" gorm:"not null;default:waiting approval"`

	DriverID      uuid.UUID `json:"driverId,omitempty" gorm:"type:uuid;default:NULL"`
	DriverTopupID uuid.UUID `json:"driverTopUpId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy     uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy     uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver      *master.Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`
	DriverTopup *DriverTopup   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverTopupID;references:ID" json:"driverTopup,omitempty"`
	types.DefaultModelProperty
}

type DriverTransactionHistoryModelAction interface {
	GetOneByID(id uuid.UUID) (m DriverTransactionHistory, err error)
	GetOneByEmail(email string) (m DriverTransactionHistory, err error)
	GetAllHistoryByDriverID(id uuid.UUID) (DriverTransactionHistory []DriverTransactionHistory, err error)

	InsertDriverTransactionHistory(p DriverTransactionHistory, tx *gorm.DB) (err error)
	DeleteDriverTransactionHistory(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewDriverTransactionHistoryAction(db *gorm.DB) DriverTransactionHistoryModelAction {
	return &DriverTransactionHistoryOrm{db}
}

func (o *DriverTransactionHistoryOrm) GetOneByID(id uuid.UUID) (DriverTransactionHistory DriverTransactionHistory, err error) {
	result := o.db.Model(&DriverTransactionHistory).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&DriverTransactionHistory)
	return DriverTransactionHistory, result.Error
}

func (o *DriverTransactionHistoryOrm) GetAllHistoryByDriverID(id uuid.UUID) (DriverTransactionHistory []DriverTransactionHistory, err error) {
	result := o.db.Model(&DriverTransactionHistory).
		Where("driver_id = ?", id).
		Preload("Driver").
		Preload("DriverTopup").
		Find(&DriverTransactionHistory)
	return DriverTransactionHistory, result.Error
}
func (o *DriverTransactionHistoryOrm) GetOneByEmail(email string) (m DriverTransactionHistory, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *DriverTransactionHistoryOrm) InsertDriverTransactionHistory(p DriverTransactionHistory, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *DriverTransactionHistoryOrm) DeleteDriverTransactionHistory(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&DriverTransactionHistory{}).Delete(&DriverTransactionHistory{}, id)
	return result.Error
}
