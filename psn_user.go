package models

import (
	"time"
)

type PsnUser struct {
	ID          uint32    `gorm:"primary_key" json:"-"`
	AccountId   int64     `json:"-"`
	PsnUserId   int64    `json:"psn_user_id"`
	PsnUserName string    `json:"psn_user_name"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
