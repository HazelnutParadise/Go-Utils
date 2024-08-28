package timeutil

import (
	"fmt"
	"time"
)

// TimeInZone 函數，接受一個 UTC 偏移值，返回該時區的當下時間
func TimeInZone(offsetHours int) time.Time {
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", offsetHours), offsetHours*3600)
	currentTime := time.Now().In(loc)
	return currentTime
}

const (
	// 定義常用的時間格式
	FormatDateOnly       = "2006-01-02"
	FormatTimeOnly       = "15:04:05"
	FormatDateTime       = "2006-01-02 15:04:05"
	FormatISO8601        = "2006-01-02T15:04:05Z07:00"
	FormatISO8601Compact = "20060102T150405Z0700"
	FormatRFC1123        = time.RFC1123
	FormatRFC822         = time.RFC822
)

// NowFormatted 根據指定格式返回當前時間的字串表示，預設使用 UTC+0
func NowFormatted(format string, timezoneOffset ...int) string {
	if len(timezoneOffset) > 1 {
		panic("NowFormatted: too many arguments, only one timezoneOffset can be specified")
	}

	offset := 0 // 預設使用 UTC+0
	if len(timezoneOffset) == 1 {
		offset = timezoneOffset[0]
	}

	loc := time.FixedZone("UTC", offset*3600)
	now := time.Now().In(loc)
	return now.Format(format)
}

// FormatTime 格式化指定時間
func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

// DaysBetween 函數，計算兩個日期之間的天數
func DaysBetween(startDate, endDate time.Time) int {
	// 忽略時間部分，只考慮日期
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	// 計算時間差
	duration := endDate.Sub(startDate)

	// 將時間差轉換為天數
	return int(duration.Hours() / 24)
}
