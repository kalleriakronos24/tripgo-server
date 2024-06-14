package base

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/booklap-be/models/system"
	"github.com/kalleriakronos24/booklap-be/types"
)

// type moduleOrm struct {
// 	db *gorm.DB
// }

type BaseSubModulePermission struct {
	ID           uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	SubModuleID  uuid.UUID `json:"subModuleId" gorm:"type:uuid;default:not null"`
	PermissionID uuid.UUID `json:"permissionId" gorm:"type:uuid;default:not null"`

	// relations
	Module     *system.SubModule   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:SubModuleID;references:ID" json:"subModule"`
	Permission *system.Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:PermissionID;references:ID" json:"permission"`
	types.DefaultModelProperty
}
