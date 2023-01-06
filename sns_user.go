package models

import (
	"time"
)

type SnsUser struct {
	ID          uint32    `gorm:"primary_key" json:"-"`
	AccountId   int64     `json:"-"`
	SnsId       int       `json:"sns_id"`
	SnsUserId   string    `json:"sns_user_id"`
	SnsUserName string    `json:"sns_user_name"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
