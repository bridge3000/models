package models

import (
	"time"
)

type AdminLog struct {
	Id        uint
	AdminId   uint
	AdminName string
	Log_type  int //日志类型，详情在AdminLog::getTypes方法里
	Message   string
	CreatedAt time.Time
	IP        string
}

func (this *AdminLog) GetTypes() map[int]string {
	return map[int]string{1: "公告", 2: "邮件", 3: "激活码", 4: "大区", 5: "账号", 6: "SDK", 7: "游戏文字内容", 8: "角色", 9: "礼包码", 10: "管理员"}
}
