package models

import (
	"time"
)

const (
	RECHARGE_ORDER_STATUS_CREATED int = 0
	RECHARGE_ORDER_STATUS_PAYED   int = 1
	RECHARGE_ORDER_STATUS_SENT    int = 2
)

var RechargeOrderStatusMap = map[int]string{
	RECHARGE_ORDER_STATUS_CREATED: "已建立订单",
	RECHARGE_ORDER_STATUS_PAYED:   "已支付未发货",
	RECHARGE_ORDER_STATUS_SENT:    "已支付已发货",
}

//充值对象，Mysql ORM，充值回调以及退款
type RechargeOrder struct {
	ID                   int `gorm:"primaryKey" json:"id"` //主键，对应英雄支付的gameOrder字段
	CreatedAt            time.Time
	CurrencyCode         string  //货币代码
	SelectedCurrencyCode string  //选择的货币代码，人民币支付实际接口参数发的TWD
	Amount               float64 `json:"amount"`
	Status               int     `json:"status"`     //订单支付及发货状态 0. 未支付未发货；1.已支付未发货；2.已支付已发货；
	AccountId            int64   `json:"account_id"` //账号ID
	ServerId             int32   `json:"server_id"`
	RoleId               string  `json:"role_id"` //角色ID
	GoodsName            string
	Description          string
	Attach               string //透传参数
	PayPlatform          int32  //支付平台 1苹果 2谷歌
	GameId               int32
	Msg                  string
	SuccessTime          time.Time
	PlatformOutTradeNo   string `json:"platform_out_trade_no"` //支付平台订单号
	Receipt              string //支付平台收据
	ProductId            string
	ShopId               int32 //商店ID
	ShopItemId           int64 //商店物品ID
}
