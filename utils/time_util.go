package utils

import (
	"time"
)

const LOC = "Asia/Chongqing"
const YMDHIS_COMMON_FORMAT = "2006-01-02 15:04:05"

type TimeUtil struct {
}

//获取毫秒时间戳
func (this *TimeUtil) GetMillsTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取指定日期的0点时间戳
//@param diffDays 据今天的天数
func (this *TimeUtil) GetDayTtp(diffDays int64) int64 {
	loc, _ := time.LoadLocation(LOC)
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	yesterdayTtp := t.Unix() + diffDays*24*3600
	return yesterdayTtp
}

//日期字符串转linux时间戳
// @param timeLayout 日期字符串格式
// @param 待转的日期字符串
// @return linux时间戳
func (this *TimeUtil) DateFormat2TimeStamp(timeLayout string, toBeCharge string) int64 {
	loc, _ := time.LoadLocation(LOC)                                //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	return sr
}

//把golang默认timedate format转换成Y-m-d H:i:s
func (this *TimeUtil) ConvertDatetimeFormat(date interface{}) string {
	dateStr := ""

	value, ok := date.(string)
	if ok {
		dateStr = value
	} else {
		value, ok := date.(time.Time)
		if ok {
			dateStr = value.Format(time.RFC3339)
		}
	}

	time := this.DateFormat2TimeStamp(time.RFC3339, dateStr)
	return this.GetDateFromInt64UnixTime(time, YMDHIS_COMMON_FORMAT)
}

//时间戳转日期字符串
func (this *TimeUtil) GetDateFromInt64UnixTime(timestamp int64, format string) string { //判断日期是否空
	if timestamp == 0 {
		return ""
	} else {
		tm := time.Unix(timestamp, 0)
		return tm.Format(format)
	}
}
