package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"size:20"`
	Email    string `gorm:"size:20"`
	Age      int    `gorm:"size:2"`
	Division string `gorm:"size:20"`
}
