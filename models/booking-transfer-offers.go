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

type BookingTransferOffersOrm struct {
	db *gorm.DB
}

type BookingTransferOffers struct {
	ID                uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	AdultSeater       int       `json:"adultSeater,omitempty" gorm:"not null"`
	ChildSeater       int       `json:"childSeater,omitempty" gorm:"not null"`
	FromLatCoordinate float32   `json:"fromLatCoordinate,omitempty" gorm:"not null"`
	FromLngCoordinate float32   `json:"fromLngCoordinate,omitempty" gorm:"not null"`
	ToLatCoordinate   float32   `json:"toLatCoordinate,omitempty" gorm:"not null"`
	ToLngCoordinate   float32   `json:"toLngCoordinate,omitempty" gorm:"not null"`
	FromLocation      string    `json:"fromLocation,omitempty" gorm:"not null"`
	ToLocation        string    `json:"toLocation,omitempty" gorm:"not null"`
	PassengerNotes    string    `json:"passengerNotes,omitempty" gorm:"default:NULL"`
	PickUpDate        time.Time `json:"pickUpDate,omitempty" gorm:"not null"`
	Price             float32   `json:"price,omitempty" gorm:"not null"`

	CustomerID uuid.UUID `json:"customerId,omitempty" gorm:"type:uuid;default:NULL"`
	CarModelID uuid.UUID `json:"carModelId,omitempty" gorm:"type:uuid;default:NULL"`

	CarModel        *master.CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarModelID;references:ID" json:"carModel,omitempty"`
	BookingTransfer *BookingTransfer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:BookingTransferID;references:ID" json:"bookingTransfer,omitempty"`
	types.DefaultModelProperty
}

type BookingTransferOffersModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransferOffers, err error)
	GetOneByEmail(email string) (m BookingTransferOffers, err error)

	InsertBookingTransferOffers(p BookingTransferOffers) (err error)
	DeleteBookingTransferOffers(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBookingTransferOffersAction(db *gorm.DB) BookingTransferOffersModelAction {
	return &BookingTransferOffersOrm{db}
}

func (o *BookingTransferOffersOrm) GetOneByID(id uuid.UUID) (BookingTransferOffers BookingTransferOffers, err error) {
	result := o.db.Model(&BookingTransferOffers).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&BookingTransferOffers)
	return BookingTransferOffers, result.Error
}

func (o *BookingTransferOffersOrm) GetOneByEmail(email string) (m BookingTransferOffers, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferOffersOrm) InsertBookingTransferOffers(p BookingTransferOffers) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *BookingTransferOffersOrm) DeleteBookingTransferOffers(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransferOffers{}).Delete(&BookingTransferOffers{}, id)
	return result.Error
}
