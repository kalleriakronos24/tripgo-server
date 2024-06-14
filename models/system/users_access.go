package system

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/booklap-be/models/master"
	"github.com/kalleriakronos24/booklap-be/types"
)

// type moduleOrm struct {
// 	db *gorm.DB
// }

type UserAccess struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name         string    `json:"name" gorm:"not null"`
	ModuleID     uuid.UUID `json:"moduleId" gorm:"type:uuid;default:not null"`
	SubModuleID  uuid.UUID `json:"subModuleId" gorm:"type:uuid;default:NULL"`
	PermissionID uuid.UUID `json:"permissionId" gorm:"type:uuid;default:not null"`
	UserID       uuid.UUID `json:"userId" gorm:"type:uuid;default:not null"`

	// relations
	Module     *Module      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"module"`
	SubModule  *SubModule   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"subModule"`
	Permission *Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"permission"`
	User       *master.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:ModuleID;references:ID" json:"user"`
	types.DefaultModelProperty
}
