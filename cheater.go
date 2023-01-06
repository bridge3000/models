package models

import (
	"time"
)

//外挂
type Cheater struct {
	ID        int `gorm:"primary_key"`
	GameId    int
	AccountId int64
	RoleId    string
	Ip        string
	CreatedAt time.Time
}
