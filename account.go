package models

import (
	"strconv"
	"strings"
	"time"
)

const (
	ACCOUNT_STATUS_DEFAULT int = 0
	ACCOUNT_STATUS_SUCCESS int = 1
	ACCOUNT_STATUS_BANNED  int = -1
)

type Account struct {
	Uid           int64 `gorm:"primary_key"`
	Status        int   //1正常 -1封禁
	LastLoginIp   string
	LastLoginTime time.Time
	CreatedAt     time.Time
	ChannelId     int
	DeviceId      string
	Email         string
	Password      string
	//	LastLoginSdkId int    //上次登录的社交SDKID，客户端需要
	LastLoginSdkId int    //上次登录的社交SDKID，客户端需要
	LockTime       int64  //被封禁后解封的时间戳
	RegionId       string `json:"region_id"`
	LoginedGameIds string //多个game_id使用|分割
	//	NewsFeed               int    //来源
	BirthDate   time.Time
	NickName    string
	DisplayName string //登录后的显示名称
	//	UserLoginDays          int32  `gorm:"-"` //已廢棄，暫時保留一個版本，防止客戶端報錯
	SdkCumulativeDaysLogin int32  //SDK（账号）累计登陆天数
	SdkContinuousDaysLogin int32  //SDK（账号）连续登陆天数
	SdkInfo                string //需要上报给数数的属性
	ServerId               int
}

type SdkInfo struct {
	//	"#device_id":      string `json:"#device_id"`
	//	"#carrier":      string `json:"#carrier"`
	//	"#os":       string `json:"#os"`
	//	"#screen_height":1200,      int `json:"#screen_height"`
	//	"#screen_width":1920,      int `json:"region_id"`
	//	"#manufacturer":"NVIDIA Corporation",      string `json:"#manufacturer"`
	//	"#device_model":"System Product Name (System manufacturer)",      string `json:"#device_model"`
	//	"#system_language":"zh",      string `json:"#system_language"`
	//	"#os_version":"Windows 10  (10.0.17763) 64bit",      string `json:"#os_version"`
	//	"#network_type":"LAN",      string `json:"#network_type"`
	//	"#bundle_id":"com.sgame.phantombladeex",      string `json:"#bundle_id"`
	//	"#app_version":"1.0.0",      string `json:"#app_version"`
	//	"#zone_offset":8.0      float64 `json:"#zone_offset"`
	AccountId string `json:"account_id"`
}

func GetNewsFeeds() map[int]string {
	return map[int]string{
		1: "Facebook",
		2: "Twitter",
		3: "Naver Cafe",
		4: "TapTap",
		5: "Others",
	}
}

func (this *Account) Ban(lockTime int64) {
	this.LockTime = lockTime
	this.Status = ACCOUNT_STATUS_BANNED
}

//解封
func (this *Account) Free() {
	this.LockTime = 0
	this.Status = ACCOUNT_STATUS_SUCCESS
}

//登录事件，累加登錄次數
//@return 是否登入被封禁
func (this *Account) OnLogined(ip string, lastLoginSdkId int) bool {
	nowTime := time.Now()
	format := "2006-01-02"

	if this.LockTime > nowTime.Unix() {
		this.Status = ACCOUNT_STATUS_BANNED
		return false
	} else {
		if this.LastLoginTime.Format(format) == time.Unix(nowTime.Unix()-24*3600, 0).Format(format) { //上次登录是昨天
			this.SdkContinuousDaysLogin++
			this.SdkCumulativeDaysLogin++
		} else if this.LastLoginTime.Format(format) == nowTime.Format(format) { //上次登录是今天

		} else { //更早
			this.SdkContinuousDaysLogin = 1
			this.SdkCumulativeDaysLogin++
		}

		this.Status = ACCOUNT_STATUS_SUCCESS
		this.LastLoginTime = nowTime
		this.LastLoginIp = ip
		//		this.LastLoginSdkId = lastLoginSdkId

		return true
	}
}

//借用账号存REDIS的KEY
func (this *Account) GetBorrowRedisKey() string {
	return "gmt:account_borrow_" + strconv.Itoa(int(this.Uid))
}

func (this *Account) AddNewLoginedGameId(gameId int32) {
	arr := this.getLoginedGameIdArray()
	arr = append(arr, strconv.Itoa(int(gameId)))
	this.LoginedGameIds = strings.Join(arr, "|")
}

func (this *Account) getLoginedGameIdArray() []string {
	arr := []string{}

	if this.LoginedGameIds != "" {
		arr = strings.Split(this.LoginedGameIds, "|")
	}
	return arr
}

/* 是否在本游戏的首次登录 */
func (this *Account) IsFirstLoginGame(gameId int32) bool {
	arr := this.getLoginedGameIdArray()
	strGameId := strconv.Itoa(int(gameId))

	for _, element := range arr {
		if strGameId == element {
			return false
		}
	}

	return true
}
