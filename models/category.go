package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          string  `json:"ID" gorm:"not null;type:char(36);primaryKey"`
	Name        string  `json:"name,omitempty" gorm:"not null;type:varchar(255);"`
	Description *string `json:"description,omitempty" gorm:"not null;type:varchar(255);"`
}
