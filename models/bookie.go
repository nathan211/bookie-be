package models

import "time"

type Bookie struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    BookedTime  time.Time `json:"booked_time"`
    ArrivalTime time.Time `json:"arrival_time"`
    Status      string    `json:"status"`
    UserID      uint      `json:"user_id"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
}