package models

import (
	"time"
)

type WebSource struct {
	ID        int    `gorm:"primary_key"`
	Name      string `json:"name"`
	CreatedAt time.Time
	Suffix    string //网页后缀
	Uv        int
	SignupCnt int //预约数
}
