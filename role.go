package models

import (
//	"strconv"
)

//玩家角色
type Role struct {
	RoleId        int    `json:"roleid"`
	RoleName      string `json:"rolename"`
	AccountId     int    `json:"account_id"`
	BanType       int    `json:"ban_type"`
	BanReason     string `json:"banreason"`
	BanTime       int    `json:"bantime"`
	Diamond       int    `json:"diamond"`
	LastLoginTime int    `json:"lastlogin_time"`
	Level         int    `json:"level"`
	RegistTime    int    `json:"register_time"`
	ServerId      string `json:"serverid"`
	TotalPay      int    `json:"totalpay"`
	VipLevel      int    `json:"viplevel"`
	HeroId        int    `json:"hero_id"` //1魂 2沐小葵 3弦者 4左殇 5小厮 6虚空
	GsId          int64  `json:"gs_id"`
}

func (this *Role) GetRoleServerId() int {
	return this.RoleId / ZONE_PLAYER_RATE
}

func (this *Role) GetHeroName() string {
	heros := []string{"", "魂", "沐小葵", "弦者", "左殇", "小厮", "虚空"}
	return heros[this.HeroId]
}
