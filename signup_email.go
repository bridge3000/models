package models

import (
	"time"
)

type SignupEmail struct {
	ID        int
	Email     string
	Lang      string
	Source    string
	Os        string
	CreatedAt time.Time
	IsNew     int
	Suffix    string
}
