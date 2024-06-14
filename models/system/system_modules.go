package system

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/booklap-be/types"
)

// type moduleOrm struct {
// 	db *gorm.DB
// }

type BaseModulePermission struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	ModuleID     uuid.UUID `json:"moduleId" gorm:"type:uuid;default:not null"`
	PermissionID uuid.UUID `json:"permissionId" gorm:"type:uuid;default:not null"`

	// relations
	Module     *Module      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"module"`
	Permission *Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"permission"`
	types.DefaultModelProperty
}

type Module struct {
	ID   uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name string    `json:"name" gorm:"not null"`

	// populate relation
	SubModule            []*SubModule            `json:"subModule,omitempty"`
	BaseModulePermission []*BaseModulePermission `json:"baseModulePermission,omitempty"`
	types.DefaultModelProperty
}
