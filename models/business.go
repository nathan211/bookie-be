package models

type Business struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"size:100"`
	Address string `json:"address" gorm:"size:500"`
	// Lat        float `json:"lat" gorm:"size:100"`
	// Lng        float `json:"lng" gorm:"size:100"`
	BusinessTypeID uint `json:"business_type_id" gorm:"not null"`
	UserID         uint `json:"user_id" gorm:"not null"`
}
