package models

import ()

type RechargeLog struct {
	Time         string `json:"time"`
	Method       string `json:"method"`
	ClientIP     string `json:"client_ip"`
	UserAgent    string `json:"user_agent"`
	ErrorMessage string `json:"error_message"`
	TransferType string `json:"transfer_type"`
	Action       string
	OrderNo      string `json:"order_no"`
	Ip           string
	UserId       string `json:"user_id"`
	ChannelId    string `json:"channel_id"`
	ZoneId       string `json:"zone_id"`
	RoleId       string `json:"role_id"`
	ShopId       string `json:"shop_id"`
	ShopItemId   string `json:"shop_item_id"`
	Amount       string `json:"amount"`
	HeroId       string `json:"hero_id"`
	Imei         string `json:"imei"`
	AndroidId    string `json:"android_id"`
	Oaid         string `json:"oaid"`
	RyOs         string `json:"ry_os"`
	RyDeviceType string `json:"ry_device_type"`
}
