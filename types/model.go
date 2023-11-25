package types

import (
	"gorm.io/gorm"
	"time"
)

type DefaultModelProperty struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
