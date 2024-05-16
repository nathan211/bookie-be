package models

import (
	"time"

	"gorm.io/gorm"
)

type Bookie struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	BookedTime  time.Time
	ArrivalTime time.Time
	Status      string `gorm:"size:20"`
	UserID      uint
	User        User
}