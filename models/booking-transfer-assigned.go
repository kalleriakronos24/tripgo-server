package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingTransferAssignedOrm struct {
	db *gorm.DB
}

type BookingTransferAssigned struct {
	ID    uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Price string    `gorm:"not null" json:"price,omitempty" binding:"required"`

	BookingTransferID uuid.UUID `json:"bookingTransferid,omitempty" gorm:"type:uuid;default:NULL"`
	DriverID          uuid.UUID `json:"driverid,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy         uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy         uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	BookingTransfer *BookingTransfer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:BookingTransferID;references:ID" json:"bookingTransfer,omitempty"`
	Driver          *master.Driver   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driverId,omitempty"`

	types.DefaultModelProperty
}

type BookingTransferAssignedModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransferAssigned, err error)
	GetOneByEmail(email string) (m BookingTransferAssigned, err error)

	InsertBookingTransferAssigned(p BookingTransferAssigned) (err error)
	DeleteBookingTransferAssigned(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBookingTransferAssignedAction(db *gorm.DB) BookingTransferAssignedModelAction {
	return &BookingTransferAssignedOrm{db}
}

func (o *BookingTransferAssignedOrm) GetOneByID(id uuid.UUID) (BookingTransferAssigned BookingTransferAssigned, err error) {
	result := o.db.Model(&BookingTransferAssigned).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&BookingTransferAssigned)
	return BookingTransferAssigned, result.Error
}

func (o *BookingTransferAssignedOrm) GetOneByEmail(email string) (m BookingTransferAssigned, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferAssignedOrm) InsertBookingTransferAssigned(p BookingTransferAssigned) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *BookingTransferAssignedOrm) DeleteBookingTransferAssigned(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransferAssigned{}).Delete(&BookingTransferAssigned{}, id)
	return result.Error
}
