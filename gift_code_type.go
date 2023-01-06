package models

import (
//	"../utils"
	//	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
//	"strings"
	"time"
)

type GiftCodeType struct {
	gorm.Model
	Name                 string
	DropId               int32
	StartedAt            time.Time
	ExpiredAt            time.Time
	ChannelId            string
	GiftCodeCnt          int64
	ExchangedCnt         int
	AccountExchangeLimit int64  //此类型每账号可兑换次数 0表示不限制
	RoleExchangeLimit    int64  //此类型每角色可兑换次数 0表示不限制
	ExchangeType         int    //1普通码 2通用码
	ExchangeLimit        int64  //单码可使用次数，手动添加的类型默认值是1
	ProjectId            string //项目ID 英雄USDK登录后传入
	IsLocked             int
}

func (GiftCodeType) TableName() string {
	return "gift_code_types"
}

func (this *GiftCodeType) CheckChannel(channelId string, projectId string) bool {
	if this.ChannelId == "" { //不限制渠道
		return true
	} else {
		println("exchange debug:", this.ProjectId, projectId)

//		arr := strings.Split(this.ChannelId, ",")
//		arrayUtil := utils.ArrayUtil{}
		//if arrayUtil.InStringArray(channelId, arr, true) {
			//if this.ProjectId != "" && arrayUtil.InStringArray(channelId, []string{"18", "56"}, false) { //官渠约束ProjectId的
				//return arrayUtil.InStringArray(projectId, strings.Split(this.ProjectId, ","), true)
			//} else {
				//return true
			//}
		//} else {
			//return false
		//}
		
		return true
	}
}
