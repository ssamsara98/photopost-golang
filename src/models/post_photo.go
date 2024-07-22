package models

import (
	"time"

	"gorm.io/gorm"
)

// Post model
type PostPhoto struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	ID        *string        `json:"id" gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Keypath   string         `json:"keypath" gorm:"not null"`
}
