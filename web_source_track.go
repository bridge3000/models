package models

import (
	"time"
)

type WebSourceTrack struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	WebSourceId int
	Ip          string
	Ua          string
	Lang        string
}
