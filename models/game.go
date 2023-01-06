package models

import (
	"time"
)

type Game struct {
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	CreatedAt          *time.Time `json:"created_at"`
	PayNotifyUrl       string     `json:"pay_notify_url"`
	AppleIapPassword   string     `json:"apple_iap_password"` //苹果验证收据的password
	RedisHost          string     `json:"redis_host"`         //游戏唯一名redis
	RedisPort          string     `json:"redis_port"`
	RedisPwd           string     `json:"redis_pwd"`
	RedisDbName        string     `json:"redis_db_name"`
	GoogleIss          string     `json:"google_iss"`
	GooglePrivateKey   string     `json:"google_private_key"`
	MycardFacServiceId string     `json:"mycard_fac_service_id"` //廠商服務代碼
	MycardKey          string     `json:"mycard_key"`            //廠商的 Key
	TaAppId            string     `json:"ta_app_id"`             //数数科技APPID
}
