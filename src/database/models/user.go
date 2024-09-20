package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"unique;not null" json:"nickname"`
	Email    string `gorm:"unique;not null" json:"email"`
	Bio      string `gorm:"type:text" json:"bio"`
	Password string `gorm:"not null" json:"password"`
}
