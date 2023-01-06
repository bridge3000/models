package models

import (
	"time"
)

type Channel struct {
	ID                  int       `gorm:"primary_key" json:"-"`
	ChannelId           int       `json:"channel_id"`
	Name                string    `json:"name"`
	AppUrl              string    `json:"app_url"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
	RechargeEnable      bool      `json:"recharge_enable"`
	GuestLoginEnable    bool      `json:"guest_login_enable"`
	EmailLoginEnable    bool      `json:"email_login_enable"`
	FacebookLoginEnable bool      `json:"facebook_login_enable"`
	GoogleLoginEnable   bool      `json:"google_login_enable"`
	TwitterLoginEnable  bool      `json:"twitter_login_enable"`
	AppleLoginEnable    bool      `json:"apple_login_enable"`
	TapLoginEnable      bool      `json:"tap_login_enable"`
	RegEnable           bool      `json:"reg_enable"`
	GiftCodeEnable      bool      `json:"gift_code_enable"`
	WindowEnable        bool      `json:"window_enable"`
	WindowContent       string    `json:"window_content"`
}
