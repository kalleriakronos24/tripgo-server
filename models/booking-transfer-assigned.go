package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
)

type BookingTransferAssignedOrm struct {
	db *gorm.DB
}

type BookingTransferAssigned struct {
	ID uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`

	BookingTransferID uuid.UUID `json:"bookingTransferid,omitempty" gorm:"type:uuid;default:NULL"`
	CarManagementID   uuid.UUID `json:"carManagementId,omitempty" gorm:"type:uuid;default:NULL"`
	IsAccepted        *bool     `json:"isAccepted" gorm:"type:boolean;default:false"`
	IsCancelled       *bool     `json:"isCancelled" gorm:"type:boolean;default:false"`
	IsOnGoing         *bool     `json:"isOngoing" gorm:"type:boolean;default:false"`
	IsPickedUp        *bool     `json:"isPickedUp" gorm:"type:boolean;default:false"`
	IsCompleted       *bool     `json:"isCompleted" gorm:"type:boolean;default:false"`
	CreatedBy         uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy         uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	BookingTransfer *BookingTransfer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:BookingTransferID;references:ID" json:"bookingTransfer,omitempty"`
	CarManagement   *master.CarManagement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarManagementID;references:ID" json:"carManagement,omitempty"`

	BookingTransferRating *BookingTransferRating `json:"bookingTransferRating,omitempty"`
	types.DefaultModelProperty
}

type BookingTransferAssignedModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransferAssigned, err error)
	GetOneByEmail(email string) (m BookingTransferAssigned, err error)
	GetBookingAssignedNotAcceptedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error)
	GetBookingAssignedAcceptedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error)
	GetBookingAssignedCancelledByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error)
	GetBookingAssignedOnGoingByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error)
	GetBookingAssignedCompletedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error)
	InsertBookingTransferAssigned(p BookingTransferAssigned, tx *gorm.DB) (err error)
	UpdateBookingTransferAssigned(id uuid.UUID, p BookingTransferAssigned, tx *gorm.DB) (err error)
	DeleteBookingTransferAssigned(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBookingTransferAssignedAction(db *gorm.DB) BookingTransferAssignedModelAction {
	return &BookingTransferAssignedOrm{db}
}

func (o *BookingTransferAssignedOrm) GetOneByID(id uuid.UUID) (BookingTransferAssigned BookingTransferAssigned, err error) {
	result := o.db.Model(&BookingTransferAssigned).
		Where("id = ?", id).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Preload("CarManagement", func(dbx *gorm.DB) *gorm.DB {
			return dbx.Preload("Driver", func(dbxx *gorm.DB) *gorm.DB {
				return dbxx.Preload("Credentials")
			})
		}).
		First(&BookingTransferAssigned)
	return BookingTransferAssigned, result.Error
}

func (o *BookingTransferAssignedOrm) GetOneByEmail(email string) (m BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) GetBookingAssignedNotAcceptedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("is_accepted = ? AND is_cancelled = ? AND is_on_going = ? AND is_picked_up = ? AND is_completed = ?", false, false, false, false, false).
		Preload("CarManagement", func(db *gorm.DB) *gorm.DB {
			return db.Where("driver_id", id).Preload("CarModel").Find(&master.CarManagement{})
		}).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Order("created_at DESC").
		Find(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) GetBookingAssignedAcceptedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("is_accepted = ? AND is_cancelled = ? AND is_on_going = ? AND is_picked_up = ? AND is_completed = ?", true, false, false, false, false).
		Preload("CarManagement", func(db *gorm.DB) *gorm.DB {
			return db.Where("driver_id", id).Preload("CarModel").Find(&master.CarManagement{})
		}).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Order("created_at DESC").
		Find(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) GetBookingAssignedCancelledByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("is_accepted = ? AND is_cancelled = ? AND is_on_going = ? AND is_picked_up = ? AND is_completed = ?", false, true, false, false, false).
		Preload("CarManagement", func(db *gorm.DB) *gorm.DB {
			return db.Where("driver_id", id).Preload("CarModel").Find(&master.CarManagement{})
		}).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Order("created_at DESC").
		Find(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) GetBookingAssignedOnGoingByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("is_accepted = ? AND is_cancelled = ? AND (is_on_going = ? OR is_picked_up = ?) AND is_completed = ?", false, false, true, true, false).
		Preload("CarManagement", func(db *gorm.DB) *gorm.DB {
			return db.Where("driver_id", id).Preload("CarModel").Find(&master.CarManagement{})
		}).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Order("created_at DESC").
		Find(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) GetBookingAssignedCompletedByDriverID(id uuid.UUID) (m []BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("is_accepted = ? AND is_cancelled = ? AND is_on_going = ? AND is_picked_up = ? AND is_completed = ?", false, false, false, false, true).
		Preload("CarManagement", func(db *gorm.DB) *gorm.DB {
			return db.Where("driver_id", id).Preload("CarModel").Find(&master.CarManagement{})
		}).
		Preload("BookingTransfer", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Customer", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Credentials")
			})
		}).
		Order("created_at DESC").
		Find(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) InsertBookingTransferAssigned(p BookingTransferAssigned, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}
func (o *BookingTransferAssignedOrm) UpdateBookingTransferAssigned(id uuid.UUID, p BookingTransferAssigned, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *BookingTransferAssignedOrm) DeleteBookingTransferAssigned(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransferAssigned{}).Delete(&BookingTransferAssigned{}, id)
	return result.Error
}
