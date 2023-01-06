package models

import (
	"time"
)

type VistorUser struct {
	Vid            int64 `gorm:"primary_key"`
	VistorAccount  string
	VistorPassword string
	CreatedAt      time.Time
	AccountId      int64
}
