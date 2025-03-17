package model

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID            int       `gorm:"primary_key;column:id"`
	SavedName     string    `gorm:"column:saved_name"`
	UserID        int       `gorm:"column:user_id"`
	ContactUserID int       `gorm:"column:contact_user_id"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
