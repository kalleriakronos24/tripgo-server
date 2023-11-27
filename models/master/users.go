package models

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userOrm struct {
	db *gorm.DB
}

type User struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name     string    `json:"name" gorm:"not null" binding:"required"`
	Address  string    `json:",omitempty"`
	Username string    `json:",omitempty" binding:"required" gorm:"not null"`
	Email    string    `gorm:"email:id,unique" json:",omitempty" binding:"required" gorm:"not null"`
	Password string    `json:",omitempty" binding:"required" gorm:"not null"`
	Role     string    `json:",omitempty" binding:"required" gorm:"not null"`

	CreatedBy           uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy           uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`
	CreatedBySuperAdmin *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CreatedBy;references:ID" json:"createdBySuperAdmin"`
	UpdatedBySuperAdmin *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UpdatedBy;references:ID" json:"updatedBySuperAdmin"`

	types.DefaultModelProperty
}

type UserModelAction interface {
	GetOneByID(id uuid.UUID) (m User, err error)
	GetOneByUserName(username string) (m User, err error)
	GetOneByEmail(email string) (m User, err error)

	InsertUser(p User) (err error)

	UpdateUser(id uuid.UUID, p User) (err error)
}

func NewUserAction(db *gorm.DB) UserModelAction {
	return &userOrm{db}
}

func (o *userOrm) GetOneByID(id uuid.UUID) (user User, err error) {

	result := o.db.Model(&user).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByEmail(email string) (m User, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *userOrm) GetOneByUserName(username string) (m User, err error) {
	result := o.db.Model(&m).Where("username = ?", username).First(&m)
	return m, result.Error
}

func (o *userOrm) InsertUser(p User) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *userOrm) UpdateUser(id uuid.UUID, p User) (err error) {
	result := o.db.Model(&p).Where("id", id).Updates(&p)
	return result.Error
}
