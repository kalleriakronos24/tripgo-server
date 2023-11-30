package models

import (
	"github.com/google/uuid"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type DocumentOrm struct {
	db *gorm.DB
}

type Document struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Base64       string    `json:"base64,omitempty"`
	Path         string    `json:"path,omitempty" gorm:"not null"`
	AbsolutePath string    `json:"absolutePath,omitempty" gorm:"not null"`
	FileName     string    `json:"filename,omitempty" gorm:"not null"`
	Extension    string    `json:"extension,omitempty" gorm:"not null;"`
	Location     string    `json:"location,omitempty" gorm:"not null;default:local"`

	DocumentCreatedBy uuid.UUID          `json:"createdBy" gorm:"type:uuid;not null;default:NULL;"`
	DocumentUpdatedBy uuid.UUID          `json:"updatedBy" gorm:"type:uuid;default:NULL;"`
	CreatedByUser     *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DocumentCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser     *masterModels.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DocumentUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type DocumentModelAction interface {
	GetAllDocument(userId uuid.UUID) (m []Document, err error)
	GetOneDocumentByID(id uuid.UUID) (m Document, err error)
	GetOneDocumentByFileName(fileName string) (m Document, err error)
	InsertDocument(p Document) (m Document, err error)
	UpdateDocument(id uuid.UUID, p Document) (m Document, err error)
}

func NewDocumentAction(db *gorm.DB) DocumentModelAction {
	return &DocumentOrm{db}
}

func (o *DocumentOrm) GetAllDocument(userId uuid.UUID) (m []Document, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}).
				First(&masterModels.User{}, userId)
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Find(&m)
	return m, result.Error
}

func (o *DocumentOrm) GetOneDocumentByID(id uuid.UUID) (m Document, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		First(&m, id)
	return m, result.Error
}

func (o *DocumentOrm) GetOneDocumentByFileName(fileName string) (m Document, err error) {
	result := o.db.Model(&m).Where("lower(filename) = ?", strings.ToLower(fileName)).First(&m)
	return m, result.Error
}

func (o *DocumentOrm) InsertDocument(p Document) (m Document, err error) {
	result := o.db.Model(&p).Omit(clause.Associations).Create(&p)
	return p, result.Error
}

func (o *DocumentOrm) UpdateDocument(id uuid.UUID, p Document) (m Document, err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return p, result.Error
}
