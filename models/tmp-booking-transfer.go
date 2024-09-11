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

type TempBookingTransferOrm struct {
	db *gorm.DB
}

type TempBookingTransfer struct {
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

	CustomerID uuid.UUID `json:"customerId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy  uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy  uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Customer *master.Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CustomerID;references:ID" json:"customer,omitempty"`

	types.DefaultModelProperty
}

type TempBookingTransferModelAction interface {
	GetOneByID(id uuid.UUID) (m TempBookingTransfer, err error)
	GetOneByEmail(email string) (m TempBookingTransfer, err error)

	InsertTempBookingTransfer(p TempBookingTransfer) (err error)
	DeleteTempBookingTransfer(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewTempBookingTransferAction(db *gorm.DB) TempBookingTransferModelAction {
	return &TempBookingTransferOrm{db}
}

func (o *TempBookingTransferOrm) GetOneByID(id uuid.UUID) (TempBookingTransfer TempBookingTransfer, err error) {
	result := o.db.Model(&TempBookingTransfer).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&TempBookingTransfer)
	return TempBookingTransfer, result.Error
}

func (o *TempBookingTransferOrm) GetOneByEmail(email string) (m TempBookingTransfer, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *TempBookingTransferOrm) InsertTempBookingTransfer(p TempBookingTransfer) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *TempBookingTransferOrm) DeleteTempBookingTransfer(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&TempBookingTransfer{}).Delete(&TempBookingTransfer{}, id)
	return result.Error
}
