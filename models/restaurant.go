package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID   int64 `gorm:"primaryKey"`
	Name string
}
