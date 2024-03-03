package models

import "gorm.io/gorm"

type MenuItem struct {
	gorm.Model
	ID           int64 `gorm:"primaryKey"`
	Name         string
	RestaurantID int64
}
