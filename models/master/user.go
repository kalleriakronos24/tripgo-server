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
	Username string    `gorm:"uniqueIndex" json:"-"`
	Email    string    `gorm:"unique" json:"-"`
	Password string    `json:"-"`
	Bio      string    `json:"-"`
	Role     string    `json:"-"`

	types.DefaultModelProperty
}

type UserModelAction interface {
	GetOneByID(id uuid.UUID) (user User, err error)
	GetOneByUsername(username string) (user User, err error)
	InsertUser(user User) (id uuid.UUID, err error)
	UpdateUser(user User) (err error)
}

func NewUserAction(db *gorm.DB) UserModelAction {
	return &userOrm{db}
}

func (o *userOrm) GetOneByID(id uuid.UUID) (user User, err error) {
	result := o.db.Model(&User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByUsername(username string) (user User, err error) {
	result := o.db.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

func (o *userOrm) InsertUser(user User) (id uuid.UUID, err error) {
	result := o.db.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func (o *userOrm) UpdateUser(user User) (err error) {
	result := o.db.Model(&User{}).Model(&user).Updates(&user)
	return result.Error
}
