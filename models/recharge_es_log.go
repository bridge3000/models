package models

import (
	"../sdk_protos"
	"time"
)

//记入ES里的日志
type RechargeESLog struct {
	Code   sdk_protos.EStateCode `json:"code"`
	Action string                `json:"action"`
	//	OrderId      int                   `json:"order_id,omitempty"` //灵游坊订单号
	AccountId        int64     `json:"account_id,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	TransferType     string    `json:"transfer_type"`
	IP               string    `json:"ip"`
	GameId           int32     `json:"game_id,omitempty"`
	RoleId           string    `json:"role_id,omitempty"`
	OutTradeNo       string    `json:"out_trade_no,omitempty"` //灵游坊订单号
	ServerId         int32     `json:"server_id,omitempty"`
	Receipt          string    `json:"receipt,omitempty"`
	IapTransactionId string    `json:"iap_transaction_id,omitempty"` //支付平台订单号
	PackageName      string    `json:"package_name,omitempty"`
	ProductId        string    `json:"product_id,omitempty"`
	PayPlatform      int32     `json:"pay_platform,omitempty"`
	Attach           string    `json:"attach,omitempty"`
	Amount           float64   `json:"amount"`
	CurrencyCode     string    `json:"currency_code,omitempty"`
	Ttp              int64     `json:"ttp"`
	Sign             string    `json:"sign,omitempty"`
	StartDateTime    string    `json:"start_date_time,omitempty"`
	EndDateTime      string    `json:"end_date_time,omitempty"`
	MyCardTradeNo    string    `json:"my_card_trade_no,omitempty"`
	ErrMsg           string    `json:"err_msg,omitempty"`
	AuthCode         string    `json:"auth_code,omitempty"`
	CardId           string    `json:"card_id,omitempty"`
	CardPw           string    `json:"card_pw,omitempty"`
	AccessToken      string    `json:"access_token,omitempty"`
	DragonMark       int       `json:"dragon_mark,omitempty"`
	TransactionUrl   string    `json:"transaction_url,omitempty"`
	ShopId           int32     `json:"shop_id,omitempty"`
	ShopItemId       int64     `json:"shop_item_id,omitempty"`
}
