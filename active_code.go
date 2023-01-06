package models

import (
	"time"
)

type ActiveCode struct {
	Id         int
	ActiveCode string //游客账号
	AccountId  int64  //账号ID
	CreatedAt  time.Time
	LogId      int
}

func (ActiveCode) TableName() string {
	return "active_codes"
}
