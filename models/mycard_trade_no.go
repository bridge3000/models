package models

import (
	"time"
)

type MycardTradeNo struct {
	ID          int
	TradeNo     string
	OutTradeNo  string
	CreatedAt   time.Time
	PaymentType string //保存一下 diff的时候必须要回传给mycard
}
