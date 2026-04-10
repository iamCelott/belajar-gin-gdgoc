package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID        string `json:"ID" gorm:"type:char(36);primaryKey"`
	Name      string `json:"name" gorm:"not null;default:'';type:varchar(255)"`
	Path      string `json:"path" gorm:"not null;default:''"`
	Type      string `json:"type" gorm:"not null;default:''"`
	Extension string `json:"extension" gorm:"not null;default:''"`
	OwnerID   string `json:"owner_id" gorm:"type:uuid;not null"`
	OwnerType string `json:"owner_type" gorm:"not null;type:varchar(50)"`
}
