package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"size:100"`
	Email    string `gorm:"size:100;unique"`
	Password string `gorm:"size:100"`
}