package models

import (
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
	ID uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`

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
	GrandTotal        float32   `json:"grandTotal,omitempty" gorm:"default:NULL"`
	AddPickupPoint    int       `json:"addPickupPoint,omitempty" gorm:"default:0"`
	AddDropPoint      int       `json:"addDropPoint,omitempty" gorm:"default:0"`
	Distance          float32   `json:"distance,omitempty" gorm:"default:0"`
	Status            string    `json:"status,omitempty" gorm:"not null;default:waiting for driver accept"`
	Uid               string    `json:"uid,omitempty" gorm:"default:NULL"`

	CustomerID uuid.UUID `json:"customerId,omitempty" gorm:"type:uuid;not null"`
	CarModelID uuid.UUID `json:"carModelId,omitempty" gorm:"type:uuid;not null"`
	CreatedBy  uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy  uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	CarModel *master.CarModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarModelID;references:ID" json:"carModel,omitempty"`
	Customer *master.Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CustomerID;references:ID" json:"customer,omitempty"`

	BookingTransferAssigned *BookingTransferAssigned `json:"driverAssigned,omitempty"`
	types.DefaultModelProperty
}

type BookingTransferModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransfer, err error)
	GetOneByEmail(email string) (m BookingTransfer, err error)
	GetAllByCustomerID(id uuid.UUID) (BookingTransfer []*BookingTransfer, err error)
	GetCountByCustomerID(id uuid.UUID) (ctx int64, err error)

	InsertBookingTransfer(p BookingTransfer, tx *gorm.DB) (bookingTransfer BookingTransfer, err error)
	UpdateBookingTransfer(id uuid.UUID, p BookingTransfer, tx *gorm.DB) (err error)
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

func (o *BookingTransferOrm) GetAllByCustomerID(id uuid.UUID) (BookingTransfer []*BookingTransfer, err error) {
	result := o.db.Model(&BookingTransfer).
		Where("customer_id = ?", id).
		Preload("CarModel").
		Preload("Customer").
		Preload("BookingTransferAssigned", func(db *gorm.DB) *gorm.DB {
			return db.Preload("BookingTransferRating").Preload("CarManagement", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Driver")
			})
		}).
		Order("created_at DESC").
		Find(&BookingTransfer)
	return BookingTransfer, result.Error
}

func (o *BookingTransferOrm) GetCountByCustomerID(id uuid.UUID) (ctx int64, err error) {
	result := o.db.Model(&BookingTransfer{}).Where("customer_id = ?", id).Count(&ctx)
	return ctx, result.Error
}

func (o *BookingTransferOrm) GetCountActiveByCustomerID(id uuid.UUID) (ctx int64, err error) {
	result := o.db.Model(&BookingTransfer{}).Where("customer_id = ?", id).Preload("BookingTransferAssigned", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_completed != ? OR is_cancelled != ?", true, true).First(&BookingTransferAssigned{})
	}).Count(&ctx)
	return ctx, result.Error
}

func (o *BookingTransferOrm) GetOneByEmail(email string) (m BookingTransfer, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferOrm) InsertBookingTransfer(p BookingTransfer, tx *gorm.DB) (bookingTransfer BookingTransfer, err error) {
	result := tx.Model(&p).Create(&p)
	return p, result.Error
}

func (o *BookingTransferOrm) UpdateBookingTransfer(id uuid.UUID, p BookingTransfer, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *BookingTransferOrm) DeleteBookingTransfer(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransfer{}).Delete(&BookingTransfer{}, id)
	return result.Error
}
