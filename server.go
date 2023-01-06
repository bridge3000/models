package models

import (
	"strings"
	"time"
)

type Server struct {
	ID               int       `gorm:"primary_key" json:"id"`
	ServerId         int       `json:"server_id"`
	ServerName       string    `json:"server_name"`
	GameId           int       `json:"game_id"`
	IpAddress        string    `json:"ip_address"`
	LinkPort         int       `json:"link_port"`
	LinkCount        int       `json:"link_count"`
	CreatedAt        time.Time `json:"created_at"`
	IconName         string    `json:"icon_name"`
	IsOpen           int       `json:"is_open"`
	State            int       `json:"state"`
	Tag              int32     `json:"tag"`
	Domain           string    `json:"domain"`
	RegionIds        string    `json:"region_ids"`
	IsDefault        int       `json:"is_default"`
	OnlineCnt        int       `gorm:"-"`
	AccountIncreased int       `gorm:"-"`
}

const ZONE_PLAYER_RATE = 10000000 //大区和角色ID的倍数关系 发全服邮件

func (this *Server) GetZoneMailPlayerId() int {
	return this.ServerId * ZONE_PLAYER_RATE
}

//账号地区是否在此服务器的区域范围
func (this *Server) InRegion(accountRegionId string) bool {
	tmpArr := strings.Split(this.RegionIds, "|")

	for _, regionId := range tmpArr {
		if regionId == accountRegionId {
			return true
		}
	}
	return false
}
