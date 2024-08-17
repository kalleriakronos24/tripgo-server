package system

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/types"
)

// type moduleOrm struct {
// 	db *gorm.DB
// }

type BaseSubModulePermission struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	SubModuleID  uuid.UUID `json:"subModuleId" gorm:"type:uuid;default:not null"`
	PermissionID uuid.UUID `json:"permissionId" gorm:"type:uuid;default:not null"`

	// relations
	Module     *SubModule   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:SubModuleID;references:ID" json:"subModule"`
	Permission *Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:PermissionID;references:ID" json:"permission"`
	types.DefaultModelProperty
}

type SubModule struct {
	ID       uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name     string    `json:"name" gorm:"not null"`
	ModuleID uuid.UUID `json:"moduleId" gorm:"type:uuid;default:NULL"`

	// relations
	Module *master.Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"module"`
	// populate relation
	BaseSubModulePermission []*BaseSubModulePermission `json:"baseSubModulePermission,omitempty"`
	types.DefaultModelProperty
}
