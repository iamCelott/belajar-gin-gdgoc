package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   `json:"ID" gorm:"primaryKey;"`
	Name        string `json:"name" gorm:"not null;type:varchar(255);"`
	Category    string `json:"category" gorm:"not null;type:varchar(255);"`
	Description *string `json:"description,omitempty" gorm:"type:varchar(255);"`
}
