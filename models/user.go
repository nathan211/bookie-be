package models

type User struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	FullName   string     `json:"full_name" gorm:"size:100"`
	Email      string     `json:"email" gorm:"size:100;unique"`
	Password   string     `json:"password" gorm:"size:100"`
	Businesses []Business `gorm:"foreignKey:UserID"`
}
