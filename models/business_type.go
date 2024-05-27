package models

type BusinessType struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"size:300"`
	Businesses []Business `gorm:"foreignKey:BusinessID"`
}
