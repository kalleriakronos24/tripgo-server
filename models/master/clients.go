package master

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/types"
	"gorm.io/gorm"
	"strings"
)

type clientOrm struct {
	db *gorm.DB
}

type Client struct {
	ID          uuid.UUID `json:"id" gorm:"index:id,unique;type:uuid;default:gen_random_uuid();"`
	Name        string    `json:"name" gorm:"not null;default:NULL" `
	PhoneNumber string    `json:"phoneNumber,omitempty" gorm:"not null;default:NULL"`
	Email       string    `gorm:"index:email,unique;not null;default:NULL" json:"email,omitempty"`
	Address     string    `json:"address,omitempty" gorm:"default:NULL"`

	CompanyID uuid.UUID `json:"companyId" gorm:"type:uuid;default:NULL"`
	Company   *Company  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:CompanyID;references:ID" json:"company"`

	ClientCreatedBy uuid.UUID `json:"createdBy" gorm:"type:uuid;not null;default:NULL"`
	ClientUpdatedBy uuid.UUID `json:"updatedBy" gorm:"type:uuid;default:NULL"`
	CreatedByUser   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClientCreatedBy;references:ID" json:"createdByUser"`
	UpdatedByUser   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClientUpdatedBy;references:ID" json:"updatedByUser"`

	types.DefaultModelProperty
}

type ClientModelAction interface {
	GetAllClientPaginated(c *gin.Context, userId uuid.UUID) (pagination *database.Pagination, err error)
	GetAllClient(userId uuid.UUID) (m []Client, err error)
	GetOneClientByID(id uuid.UUID) (m Client, err error)
	GetOneClientByName(name string) (m Client, err error)
	GetOneClientByEmail(email string) (m Client, err error)
	GetOneClientByCompanyID(id uuid.UUID) (m []Client, err error)

	InsertClient(p Client) (err error)
	UpdateClient(id uuid.UUID, p Client) (err error)
	DeleteClient(id uuid.UUID, tx *gorm.DB) (err error)
	DeleteClientByCompanyID(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewClientAction(db *gorm.DB) ClientModelAction {
	return &clientOrm{db}
}

func (o *clientOrm) GetAllClientPaginated(c *gin.Context, userId uuid.UUID) (*database.Pagination, error) {

	var mArr []*Client
	var pagination database.Pagination

	user := User{
		ID: userId,
	}

	userResult := o.db.Model(&user).Where("id = ?", userId).First(&user)

	if userResult.Error != nil {
		return nil, userResult.Error
	}

	var companies []Company
	companyResult := o.db.Model(&companies).Where("company_created_by = ?", userId).Find(&companies)
	companyIds := make([]uuid.UUID, len(companies))

	if len(companies) > 0 && user.Role == "superadmin" {
		for _, company := range companies {
			companyIds = append(companyIds, company.ID)
		}
	} else {
		companyIds = append(companyIds, user.CompanyID)
	}

	if companyResult.Error != nil {
		return nil, companyResult.Error
	}

	o.db.
		Scopes(database.Paginator(c, &mArr, []string{"CreatedByUser", "UpdatedByUser", "Company"}, &pagination)).
		Where("company_id IN ?", companyIds).
		Find(&mArr)

	pagination.Data = &mArr
	return &pagination, nil
}

func (o *clientOrm) GetAllClient(userId uuid.UUID) (m []Client, err error) {

	user := User{
		ID: userId,
	}

	userResult := o.db.Model(&user).Where("id = ?", userId).First(&user)

	if userResult.Error != nil {
		return nil, userResult.Error
	}

	var companies []Company
	companyResult := o.db.Model(&companies).Where("company_created_by = ?", userId).Find(&companies)
	companyIds := make([]uuid.UUID, len(companies))

	if len(companies) > 0 && user.Role == "superadmin" {
		for _, company := range companies {
			companyIds = append(companyIds, company.ID)
		}
	} else {
		companyIds = append(companyIds, user.CompanyID)
	}

	if companyResult.Error != nil {
		return nil, companyResult.Error
	}

	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Where("company_id IN ?", companyIds).
		Find(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByID(id uuid.UUID) (m Client, err error) {
	result := o.db.Model(&m).
		Preload("CreatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"})
		}).
		Preload("UpdatedByUser", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedBy", "UpdatedAt"})
		}).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"ID", "Name", "CreatedAt", "UpdatedAt"})
		}).
		Where("id = ?", id).
		First(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByName(name string) (m Client, err error) {
	result := o.db.Model(&m).Where("lower(name) = ?", strings.ToLower(name)).First(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByCompanyID(id uuid.UUID) (m []Client, err error) {
	result := o.db.Model(&m).Where("company_id = ?", id).Find(&m)
	return m, result.Error
}

func (o *clientOrm) GetOneClientByEmail(email string) (m Client, err error) {
	result := o.db.Model(&m).Where("email = ?", email).First(&m)
	return m, result.Error
}

func (o *clientOrm) InsertClient(p Client) (err error) {
	result := o.db.Model(&p).Create(&p)
	return result.Error
}

func (o *clientOrm) UpdateClient(id uuid.UUID, p Client) (err error) {
	result := o.db.Model(&p).Where("id = ?", id).Updates(&p)
	return result.Error
}

func (o *clientOrm) DeleteClient(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Client{}).Delete(&Client{}, id)
	return result.Error
}

func (o *clientOrm) DeleteClientByCompanyID(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&Client{}).Where("company_id = ?", id).Delete(&Client{}, id)
	return result.Error
}
