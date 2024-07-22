package models

import (
	"time"

	"gorm.io/gorm"
)

// Post model
type PostToPhoto struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	ID        *string        `json:"id" gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Position  uint           `json:"position" gorm:"not null;default:0"`
	PostID    uint           `json:"postId"`
	Post      *Post          `json:"post" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PhotoID   string         `json:"photoId" gorm:"type:uuid"`
	Photo     *PostPhoto     `json:"photo" gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (ptp PostToPhoto) BeforeCreate(db *gorm.DB) error {
	// ...
	return nil
}
