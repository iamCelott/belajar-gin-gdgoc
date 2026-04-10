package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          string    `json:"ID" gorm:"not null;type:char(36);primaryKey"`
	Name        string    `json:"name,omitempty" gorm:"not null;type:varchar(255);"`
	CategoryID  string    `json:"-"`
	Category    *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Description *string   `json:"description,omitempty" gorm:"type:varchar(255);"`
	Author      string    `json:"author,omitempty" gorm:"not null;type:varchar(255);"`
	File        *File     `json:"file,omitempty" gorm:"polymorphic:Owner;"`
}
