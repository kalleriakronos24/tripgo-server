package models

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserNotificationOrm struct {
	db *gorm.DB
}

type UserNotification struct {
	ID     uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Amount float64   `gorm:"not null;default:0" json:"amount,omitempty" binding:"required"`

	DriverID  uuid.UUID `json:"driverId,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	Driver *master.Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DriverID;references:ID" json:"driver,omitempty"`

	types.DefaultModelProperty
}

type UserNotificationModelAction interface {
	GetOneByID(id uuid.UUID) (m UserNotification, err error)
	GetOneByEmail(email string) (m UserNotification, err error)
	InsertUserNotification(p UserNotification, tx *gorm.DB) (err error)
	DeleteUserNotification(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewUserNotificationAction(db *gorm.DB) UserNotificationModelAction {
	return &UserNotificationOrm{db}
}

func (o *UserNotificationOrm) GetOneByID(id uuid.UUID) (UserNotification UserNotification, err error) {
	result := o.db.Model(&UserNotification).
		Where("driver_id = ?", id).
		Preload(clause.Associations).
		First(&UserNotification)
	return UserNotification, result.Error
}

func (o *UserNotificationOrm) GetOneByEmail(email string) (m UserNotification, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *UserNotificationOrm) InsertUserNotification(p UserNotification, tx *gorm.DB) (err error) {
	result := tx.Model(&p).Create(&p)
	return result.Error
}

func (o *UserNotificationOrm) DeleteUserNotification(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&UserNotification{}).Delete(&UserNotification{}, id)
	return result.Error
}
