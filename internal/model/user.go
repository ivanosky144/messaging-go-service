package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             int            `gorm:"primary_key;column:id"`
	Username       string         `gorm:"column:username"`
	Email          string         `gorm:"column:email"`
	Password       string         `gorm:"column:password"`
	ProfilePicture string         `gorm:"column:profile_picture"`
	Desc           string         `gorm:"column:description"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Conversations  []Conversation `gorm:"foreignKey:UserID"`
	Participants   []Participant  `gorm:"foreignKey:UserID"`
	Notifications  []Notification `gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}
