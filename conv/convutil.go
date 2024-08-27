package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseF64 將任意資料轉換為 float64，錯誤時直接 panic
func ParseF64(value interface{}) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		trimmed := strings.TrimSpace(v)
		f, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			panic(fmt.Sprintf("ParseF64: cannot convert string to float64: %v", err))
		}
		return f
	default:
		panic(fmt.Sprintf("ParseF64: unsupported type: %T", value))
	}
}

// ParseF32 將任意資料轉換為 float32，錯誤時直接 panic
func ParseF32(value interface{}) float32 {
	return float32(ParseF64(value))
}

// ParseInt 將任意資料轉換為 int，錯誤時直接 panic
func ParseInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		if v > 1<<63-1 {
			panic("ParseInt: value out of range")
		}
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		trimmed := strings.TrimSpace(v)
		i, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			f, ferr := strconv.ParseFloat(trimmed, 64)
			if ferr != nil {
				panic(fmt.Sprintf("ParseInt: cannot convert string to int: %v", err))
			}
			return int(f)
		}
		return int(i)
	default:
		panic(fmt.Sprintf("ParseInt: unsupported type: %T", value))
	}
}

// ParseBool 將任意資料轉換為 bool，錯誤時直接 panic
func ParseBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		trimmed := strings.TrimSpace(strings.ToLower(v))
		if trimmed == "true" || trimmed == "1" || trimmed == "yes" || trimmed == "on" {
			return true
		} else if trimmed == "false" || trimmed == "0" || trimmed == "no" || trimmed == "off" {
			return false
		}
		num, err := strconv.ParseFloat(trimmed, 64)
		if err == nil {
			return num != 0
		}
		panic(fmt.Sprintf("ParseBool: cannot convert string to bool: %s", v))
	case int, int8, int16, int32, int64:
		return v != 0
	case uint, uint8, uint16, uint32, uint64:
		return v != 0
	case float32:
		return float32(v) != 0.0
	case float64:
		return float64(v) != 0.0
	default:
		panic(fmt.Sprintf("ParseBool: unsupported type: %T", value))
	}
}

// ToString 將任意資料轉換為字串，錯誤時直接 panic
func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
