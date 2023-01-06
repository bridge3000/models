package models

import (
	"time"
)

//礼包码
type GiftCode struct {
	Id           int
	GiftCode     string
	GiftCodeType int
	LogId        int
	//	UsedTime     int64//已记录在exchange_log里，后续版本在表里删掉字段
	CreatedAt time.Time
	//	ZoneId       int32//已记录在exchange_log里，后续版本在表里删掉字段
	//	AccountId    int32//已记录在exchange_log里，后续版本在表里删掉字段
	//	PlayerId     int32//已记录在exchange_log里，后续版本在表里删掉字段
	TypeName     string `gorm:"-"`
	ExchangeType int    `gorm:"-"` //兑换类型单一码还是通用码
}

func (GiftCode) TableName() string {
	return "gift_codes"
}
