package models

import (
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	Sort_index string `json:"-"`
	StartTime  int64
	EndTime    int64
	GameId     int
	Lang       int
}
