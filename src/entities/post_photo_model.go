package entities

import (
	"time"

	"gorm.io/gorm"
)

// Post model
type PostPhoto struct {
	ID        *string        `json:"id" gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Keypath   string         `json:"keypath" gorm:"not null"`
}
