package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string    `gorm:"size:2048"`
	Body        string    `gorm:"size:2048"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uuid.UUID `gorm:"type:uuid;default"`
}
