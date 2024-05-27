package models

type BusinessType struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"size:100"`
	Businesses []Business `gorm:"foreignKey:BusinessTypeID"`
}
