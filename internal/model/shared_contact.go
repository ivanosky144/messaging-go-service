package model

import (
	"time"

	"gorm.io/gorm"
)

type SharedContact struct {
	gorm.Model
	ID           int       `gorm:"primary_key;column:id"`
	SharedUserID string    `gorm:"column:shared_user_id"`
	MessageID    string    `gorm:"column:message_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (s *SharedContact) TableName() string {
	return "shared_contacts"
}
