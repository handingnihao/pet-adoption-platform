package utils

import (
	"time"
)

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseTime 解析时间字符串
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeStr)
}

// ParseDate 解析日期字符串
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() string {
	return FormatTime(time.Now())
}

// GetCurrentDate 获取当前日期
func GetCurrentDate() string {
	return FormatDate(time.Now())
}
