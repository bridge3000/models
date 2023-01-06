package models

import (
	"time"
)

type Content struct {
	ID        int       `json:"-"`
	Lang      int       `json:"lang"`
	ContentId int       `json:"content_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	GameId    int       `json:"game_id"`
}
