package master

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CredentialsOrm struct {
	db *gorm.DB
}

type Credentials struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Email    string    `gorm:"email:id,unique" json:",omitempty" binding:"required"`
	Password string    `json:"password,omitempty" binding:"required" gorm:"not null"`

	CreatedBy uuid.UUID `json:"createdBy,omitempty" gorm:"type:uuid;default:NULL"`
	UpdatedBy uuid.UUID `json:"updatedBy,omitempty" gorm:"type:uuid;default:NULL"`

	CredentialCustomer *Customer `json:"credentialCustomer,omitempty"`
	CredentialDriver   *Driver   `json:"credentialDriver,omitempty"`
	CredentialInternal *Internal `json:"credentialInternal,omitempty"`
	types.DefaultModelProperty
}

type CredentialsModelAction interface {
	GetOneByID(id uuid.UUID) (m Credentials, err error)
	GetOneByEmail(email string) (m Credentials, err error)

	InsertCredentials(p Credentials) (err error)
	DeleteCredentials(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewCredentialsAction(db *gorm.DB) CredentialsModelAction {
	return &CredentialsOrm{db}
}

func (o *CredentialsOrm) GetOneByID(id uuid.UUID) (Credentials Credentials, err error) {
	result := o.db.Model(&Credentials).
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&Credentials)
	return Credentials, result.Error
}

func (o *CredentialsOrm) GetOneByEmail(email string) (m Credentials, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *CredentialsOrm) InsertCredentials(p Credentials) (err error) {
	fmt.Printf("%v", p)
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *CredentialsOrm) DeleteCredentials(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Credentials{}).Delete(&Credentials{}, id)
	return result.Error
}
