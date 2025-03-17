package model

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	ID        int       `gorm:"primary_key;column:id"`
	UserID    int       `gorm:"primary_key;column:user_id"`
	MessageID int       `gorm:"primary_key;column:message_id"`
	Url       string    `gorm:"column:username"`
	FileName  string    `gorm:"column:email"`
	Size      string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *Media) TableName() string {
	return "medias"
}
