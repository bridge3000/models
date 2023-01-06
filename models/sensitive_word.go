package models

import (
	"time"
)

type SensitiveWord struct {
	Id            int
	SensitiveWord string
	ReplaceText   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
