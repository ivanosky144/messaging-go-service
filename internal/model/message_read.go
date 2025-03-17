package model

import (
	"time"

	"gorm.io/gorm"
)

type MessageRead struct {
	gorm.Model
	ID        int       `gorm:"primary_key;column:id"`
	MessageID int       `gorm:"primary_key;column:message_id"`
	UserID    int       `gorm:"primary_key;column:user_id"`
	ReadAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (m *MessageRead) TableName() string {
	return "message_reads"
}
