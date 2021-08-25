package util

import (
	"time"

	"github.com/andrezhz/go-tools/constant"
)

// ParseStr2UnixTime
func ParseStrToUnixTime(strTime string, timeLayout constant.TimeLayout) int64 {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(string(timeLayout), strTime, loc)
	return theTime.Unix()
}

// ParseUnixTime2Str
func ParseUnixTime2Str(timestamp int64, timeLayout constant.TimeLayout) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format(string(timeLayout))
}
