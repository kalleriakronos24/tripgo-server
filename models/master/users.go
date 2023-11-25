package models

import (
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
)

type userOrm struct {
	db *gorm.DB
}

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name     string    `json:"name,omitempty" binding:"required"`
	Address  string    `json:"address"`
	Username string    `json:"username,omitempty" binding:"required"`
	Email    string    `gorm:"unique" json:"email" binding:"required"`
	Password string    `json:"password,omitempty" binding:"required"`

	types.DefaultModelProperty
}

type UserModelAction interface {
	GetOneByID(id uuid.UUID) (m User, err error)
	GetOneByUserName(username string) (m User, err error)
	InsertUser(p User) (err error)
	UpdateUser(id uuid.UUID, p User) (err error)
}

func NewUserAction(db *gorm.DB) UserModelAction {
	return &userOrm{db}
}

func (o *userOrm) GetOneByID(id uuid.UUID) (user User, err error) {
	result := o.db.Model(&User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByUserName(username string) (m User, err error) {
	result := o.db.Model(&m).Where("username = ?", username).First(&m)
	return m, result.Error
}

func (o *userOrm) InsertUser(p User) (err error) {
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *userOrm) UpdateUser(id uuid.UUID, p User) (err error) {
	result := o.db.Model(&p).Where("id", id).Updates(&p)
	return result.Error
}
