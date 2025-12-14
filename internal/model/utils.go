package model

import (
	"encoding/json"
)

// ParseJSONStringArray 解析 JSON 字符串数组
func ParseJSONStringArray(jsonStr string) []string {
	if jsonStr == "" {
		return []string{}
	}

	var result []string
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return []string{}
	}
	return result
}

// ToJSONString 将字符串数组转换为 JSON 字符串
func ToJSONString(arr []string) string {
	if len(arr) == 0 {
		return "[]"
	}

	data, err := json.Marshal(arr)
	if err != nil {
		return "[]"
	}
	return string(data)
}

// ParseJSONInt64Array 解析 JSON int64 数组
func ParseJSONInt64Array(jsonStr string) []int64 {
	if jsonStr == "" {
		return []int64{}
	}

	var result []int64
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return []int64{}
	}
	return result
}

// ToJSONInt64String 将 int64 数组转换为 JSON 字符串
func ToJSONInt64String(arr []int64) string {
	if len(arr) == 0 {
		return "[]"
	}

	data, err := json.Marshal(arr)
	if err != nil {
		return "[]"
	}
	return string(data)
}
