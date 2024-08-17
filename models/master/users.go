package master

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userOrm struct {
	db *gorm.DB
}

type User struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name         string    `json:"name" gorm:"not null" binding:"required"`
	Email        string    `gorm:"email:id,unique" json:",omitempty" binding:"required"`
	Password     string    `json:"password,omitempty" binding:"required" gorm:"not null"`
	Role         string    `json:"role,omitempty" binding:"required" gorm:"not null;"`
	ProfileImage string    `json:"profileImage,omitempty" gorm:"default:NULL"`
	Phone        string    `json:"phone,omitempty" gorm:"default:NULL"`
	Status       string    `json:"status,omitempty" binding:"required" gorm:"not null;default:active;"`

	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	// populate relation
	//UserAccess []*UserAccess `json:"userAccess,omitempty"`
	types.DefaultModelProperty
}

type UserModelAction interface {
	GetOneByID(id uuid.UUID) (m User, err error)
	GetOneByUserName(username string) (m User, err error)
	GetOneByEmail(email string) (m User, err error)
	GetAllUserPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error)

	InsertUser(p User) (err error)

	UpdateUserToInactive(id uuid.UUID, tx *gorm.DB) (err error)
	RemoveUserFromCompany(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteUser(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewUserAction(db *gorm.DB) UserModelAction {
	return &userOrm{db}
}

func (o *userOrm) GetAllUserPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {
	var mArr []*User
	var pagination database.Pagination
	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedBy", "UpdatedBy", "Company"}, &pagination)).
		Where("created_by", userId).
		Find(&mArr)
	pagination.Data = &mArr
	return &pagination, nil
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

	if m.Status == "inactive" {
		return m, errors.New("user has no company related. please ask your admin for verification")
	}

	return m, result.Error
}

func (o *userOrm) UpdateUserToInactive(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&User{}).Where("created_by = ? AND role != 'superadmin'", id).Update("status", "inactive")
	return result.Error
}

func (o *userOrm) RemoveUserFromCompany(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&User{}).Where("id", id).Update("company_id", nil)
	return result.Error
}

func (o *userOrm) InsertUser(p User) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *userOrm) DeleteUser(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&User{}).Delete(&User{}, id)
	return result.Error
}

// TENTANTS
func (o *userOrm) InsertUserOwner(p User) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *userOrm) InsertUserAdmin(p User) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}
