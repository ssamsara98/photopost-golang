package models

import (
	"photopost/lib"
	"time"
)

// User model
type User struct {
	lib.ModelBase
	Email     string     `json:"email" gorm:"unique"`
	Username  string     `json:"username" gorm:"unique"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	SexType   string     `json:"sexType" gorm:"default:'unknown'"`
	Birthdate *time.Time `json:"birthdate"`
}
