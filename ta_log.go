package models

import (
	"time"
)

//数数日志
type TaLog struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	RoleId    string
	EventName string
	LogInfo   string
	Status    int
	Err       string
	GameId    int32
}
