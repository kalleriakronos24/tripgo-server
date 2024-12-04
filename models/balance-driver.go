package models

import (
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

	Driver *master.Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`

	types.DefaultModelProperty
}

type BalanceDriverModelAction interface {
	GetOneByID(id uuid.UUID) (m BalanceDriver, err error)
	GetOneByEmail(email string) (m BalanceDriver, err error)
	GetAllDriverHasEnoughBalance(price float64, carModelId uuid.UUID) (balanceDriver *BalanceDriver, err error)

	InsertBalanceDriver(p BalanceDriver, tx *gorm.DB) (err error)
	DeleteBalanceDriver(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBalanceDriverAction(db *gorm.DB) BalanceDriverModelAction {
	return &BalanceDriverOrm{db}
}

func (o *BalanceDriverOrm) GetOneByID(id uuid.UUID) (BalanceDriver BalanceDriver, err error) {
	result := o.db.Model(&BalanceDriver).
		Where("driver_id = ?", id).
		Preload(clause.Associations).
		First(&BalanceDriver)
	return BalanceDriver, result.Error
}

func (o *BalanceDriverOrm) GetAllDriverHasEnoughBalance(price float64, carModelId uuid.UUID) (balanceDriver *BalanceDriver, err error) {
	result := o.db.Model(&balanceDriver).
		Where("amount >= ? OR amount <= ?", price, price).
		Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Raw(`WITH CTE AS (
    				SELECT random() * (SELECT SUM(prob) FROM drivers) R
					)
					SELECT *
						FROM (
    						SELECT id, status, booking_status, SUM(prob) OVER (ORDER BY id) S, R
    				FROM drivers CROSS JOIN CTE
					) Q
					WHERE S >= R AND status = 'active' AND booking_status = 'ready'
					ORDER BY id
					LIMIT 1;`).First(&master.Driver{}).Preload("Credentials")
		}).First(&balanceDriver)
	return balanceDriver, result.Error
}

func (o *BalanceDriverOrm) GetOneByEmail(email string) (m BalanceDriver, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BalanceDriverOrm) InsertBalanceDriver(p BalanceDriver, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *BalanceDriverOrm) DeleteBalanceDriver(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BalanceDriver{}).Delete(&BalanceDriver{}, id)
	return result.Error
}
