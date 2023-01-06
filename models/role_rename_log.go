package models

import (
	"time"
)

//玩家改名日志
type RoleRenameLog struct {
	Id        int
	RoleId    int
	NewName   string
	AdminId   int
	AdminName string
	ServerId  int
	CreatedAt time.Time
	Result    int
	Message   string
}
