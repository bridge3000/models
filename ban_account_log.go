package models

import (
	//	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type BanAccountLog struct {
	gorm.Model
	AccountId uint64
	BanType   int //1封禁 2解封
	AdminId   uint
	AdminName string
	Reason    string
	LockTime  int64
}
