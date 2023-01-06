package models

import (
	"time"
)

type GiftCodeExchangeLog struct {
	Id           int
	GiftCode     string
	GiftCodeType int
	CreatedAt    time.Time
	ZoneId       int32
	AccountId    int32
	PlayerId     int32
}
