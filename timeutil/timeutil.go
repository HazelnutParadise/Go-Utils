package timeutil

import (
	"fmt"
	"math"
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

// DaysBetween 函數，計算兩個日期之間的天數，返回值為正數
func DaysBetween(startDate, endDate time.Time) int {
	return int(math.Abs(float64(DaysDiff(startDate, endDate))))
}

// DaysDiff 函數，計算兩個日期之間的天數，不取絕對值
func DaysDiff(startDate, endDate time.Time) int {
	// 忽略時間部分，只考慮日期
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	// 計算時間差
	duration := endDate.Sub(startDate)

	// 將時間差轉換為天數
	return int(duration.Hours() / 24)
}

// MonthsBetween 函數，計算兩個日期之間的月份數，返回值為正數
func MonthsBetween(startDate, endDate time.Time) int {
	return int(math.Abs(float64(MonthsDiff(startDate, endDate))))
}

// MonthsDiff 函數，計算兩個日期之間的月份數，不取絕對值
func MonthsDiff(startDate, endDate time.Time) int {
	// 將日期格式化為 "YYYY-MM" 的字串
	date1 := FormatTime(startDate, "2006-01")
	date2 := FormatTime(endDate, "2006-01")

	// 將字串拆分成年和月
	var year1, month1 int
	var year2, month2 int

	fmt.Sscanf(date1, "%d-%d", &year1, &month1)
	fmt.Sscanf(date2, "%d-%d", &year2, &month2)

	// 計算年和月的差異
	yearDiff := year2 - year1
	monthDiff := month2 - month1

	// 計算總月數
	totalMonths := yearDiff*12 + monthDiff

	return totalMonths
}

// YearsBetween 函數，計算兩個日期之間的年數，返回值為正數
func YearsBetween(startDate, endDate time.Time) int {
	return int(math.Abs(float64(YearsDiff(startDate, endDate))))
}

// YearsDiff 函數，計算兩個日期之間的年數，不取絕對值
func YearsDiff(startDate, endDate time.Time) int {
	// 將日期格式化為 "YYYY" 的字串
	year1 := startDate.Year()
	year2 := endDate.Year()

	// 計算年的差異
	yearDiff := year2 - year1

	return yearDiff
}
