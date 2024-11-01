package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingTransferRatingOrm struct {
	db *gorm.DB
}

type BookingTransferRating struct {
	ID uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`

	Rating                    int       `json:"rating,omitempty" gorm:"default:0"`
	Comments                  string    `json:"comments,omitempty" gorm:"default:NULL"`
	CarManagementID           uuid.UUID `json:"carManagementId,omitempty" gorm:"type:uuid;not null"`
	BookingTransferAssignedID uuid.UUID `json:"bookingTransferAssignedId,omitempty" gorm:"type:uuid;not null"`

	CarManagement           *master.CarManagement    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CarManagementID;references:ID" json:"carManagement,omitempty"`
	BookingTransferAssigned *BookingTransferAssigned `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:BookingTransferAssignedID;references:ID" json:"driverAssigned,omitempty"`

	types.DefaultModelProperty
}

type BookingTransferRatingModelAction interface {
	GetOneByID(id uuid.UUID) (m BookingTransferRating, err error)
	GetOneByEmail(email string) (m BookingTransferRating, err error)
	GetAllByCustomerID(id uuid.UUID) (BookingTransferRating []*BookingTransferRating, err error)
	GetCountByCustomerID(id uuid.UUID) (ctx int64, err error)

	InsertBookingTransferRating(p BookingTransferRating, tx *gorm.DB) (err error)
	UpdateBookingTransferRating(id uuid.UUID, p BookingTransferRating, tx *gorm.DB) (err error)
	DeleteBookingTransferRating(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewBookingTransferRatingAction(db *gorm.DB) BookingTransferRatingModelAction {
	return &BookingTransferRatingOrm{db}
}

func (o *BookingTransferRatingOrm) GetOneByID(id uuid.UUID) (BookingTransferRating BookingTransferRating, err error) {
	result := o.db.Model(&BookingTransferRating).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&BookingTransferRating)
	return BookingTransferRating, result.Error
}

func (o *BookingTransferRatingOrm) GetAllByCustomerID(id uuid.UUID) (BookingTransferRating []*BookingTransferRating, err error) {
	result := o.db.Model(&BookingTransferRating).
		Where("customer_id = ?", id).
		Preload("CarModel").
		Preload("Customer").
		Preload("BookingTransferRatingAssigned", func(db *gorm.DB) *gorm.DB {
			return db.Preload("CarManagement", func(dbx *gorm.DB) *gorm.DB {
				return dbx.Preload("Driver")
			})
		}).
		Order("created_at DESC").
		Find(&BookingTransferRating)
	return BookingTransferRating, result.Error
}

func (o *BookingTransferRatingOrm) GetCountByCustomerID(id uuid.UUID) (ctx int64, err error) {
	result := o.db.Model(&BookingTransferRating{}).Where("customer_id = ?", id).Count(&ctx)
	return ctx, result.Error
}

func (o *BookingTransferRatingOrm) GetOneByEmail(email string) (m BookingTransferRating, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *BookingTransferRatingOrm) InsertBookingTransferRating(p BookingTransferRating, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *BookingTransferRatingOrm) UpdateBookingTransferRating(id uuid.UUID, p BookingTransferRating, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *BookingTransferRatingOrm) DeleteBookingTransferRating(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&BookingTransferRating{}).Delete(&BookingTransferRating{}, id)
	return result.Error
}
