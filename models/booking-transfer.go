package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingTransferOrm struct {
	db *gorm.DB
}

type BookingTransfer struct {
	ID                    uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	FromAddress           string    `gorm:"not null" json:"fromAddress,omitempty" binding:"required"`
	DestinationAddress    string    `json:"destinationAdress,omitempty" binding:"required" gorm:"not null"`
	FromCoordinate        string    `gorm:"not null" json:"fromCoordinate,omitempty" binding:"required"`
	DestinationCoordinate string    `gorm:"not null" json:"destinationCoordinate,omitempty" binding:"required"`
	PickUpDatetime        time.Time `gorm:"not null" json:"pickUpDatetime,omitempty" binding:"required"`
	AdultSeater           int       `gorm:"not null" json:"adultSeater,omitempty" binding:"required"`
	ChildSeater           int       `gorm:"not null" json:"childSeater,omitempty" binding:"required"`
	SuggestedPrice        int       `gorm:"not null" json:"suggestedPrice,omitempty" binding:"required"`
	IsAcceptAgreement     bool      `gorm:"not null;default:FALSE;" json:"isAcceptAgreement,omitempty" binding:"required"`

	CarModelID uuid.UUID `json:"carModelId,omitempty" gorm:"type:uuid;default:NULL"`
	CustomerID uuid.UUID `json:"customerId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy  uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy  uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Customer *master.Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CustomerID;references:ID" json:"customer,omitempty"`
	CarModel *master.CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarModelID;references:ID" json:"carModel,omitempty"`

	types.DefaultModelProperty
}

type BookingTransferModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransfer, err error)
	GetOneByEmail(email string) (m BookingTransfer, err error)

	InsertBookingTransfer(p BookingTransfer) (err error)
	DeleteBookingTransfer(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBookingTransferAction(db *gorm.DB) BookingTransferModelAction {
	return &BookingTransferOrm{db}
}

func (o *BookingTransferOrm) GetOneByID(id uuid.UUID) (BookingTransfer BookingTransfer, err error) {
	result := o.db.Model(&BookingTransfer).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&BookingTransfer)
	return BookingTransfer, result.Error
}

func (o *BookingTransferOrm) GetOneByEmail(email string) (m BookingTransfer, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferOrm) InsertBookingTransfer(p BookingTransfer) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *BookingTransferOrm) DeleteBookingTransfer(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransfer{}).Delete(&BookingTransfer{}, id)
	return result.Error
}
