package system

import (
	"github.com/google/uuid"
	"github.com/kalleriakronos24/booklap-be/types"
)

// type moduleOrm struct {
// 	db *gorm.DB
// }

type Permissions struct {
	ID     uuid.UUID `gorm:"index:id,unique;type:uuid;default:gen_random_uuid();" json:"id"`
	Name   string    `json:"name" gorm:"not null"`
	Alias  string    `json:"alias" gorm:"not null"`
	Remark string    `json:"remark" gorm:"not null"`
	types.DefaultModelProperty
}
