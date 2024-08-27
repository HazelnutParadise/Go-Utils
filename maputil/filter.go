// maputil/filter.go
package maputil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/HazelnutParadise/Go-Utils/types"
)

// FilterCondition 定義篩選條件
type FilterCondition int

const (
	FilterEqualTo              FilterCondition = iota // 篩選等於目標值的鍵值對
	FilterNotEqualTo                                  // 篩選不等於目標值的鍵值對
	FilterGreaterThan                                 // 篩選大於目標值的鍵值對（僅適用於數字類型）
	FilterLessThan                                    // 篩選小於目標值的鍵值對（僅適用於數字類型）
	FilterGreaterThanOrEqualTo                        // 篩選大於或等於目標值的鍵值對（僅適用於數字類型）
	FilterLessThanOrEqualTo                           // 篩選小於或等於目標值的鍵值對（僅適用於數字類型）
	FilterContains                                    // 篩選包含目標值的字串鍵值對（僅適用於字串類型）
	FilterNotContains                                 // 篩選不包含目標值的字串鍵值對（僅適用於字串類型）
)

// FilterByKey 函數，根據鍵來篩選 map 中的鍵值對
func FilterByKey[K comparable, V any](m map[K]V, condition FilterCondition, target K) (map[K]V, error) {
	filtered := make(map[K]V)
	for k, v := range m {
		if ok, err := meetsCondition(k, condition, target); err != nil {
			return nil, err
		} else if ok {
			filtered[k] = v
		}
	}
	return filtered, nil
}

// FilterByValue 函數，根據值來篩選 map 中的鍵值對
func FilterByValue[K comparable, V comparable](m map[K]V, condition FilterCondition, target V) (map[K]V, error) {
	filtered := make(map[K]V)
	for k, v := range m {
		if ok, err := meetsCondition(v, condition, target); err != nil {
			return nil, err
		} else if ok {
			filtered[k] = v
		}
	}
	return filtered, nil
}

// CustomFilter 函數，使用自訂的篩選函數來篩選 map 中的鍵值對
func CustomFilter[K comparable, V any](m map[K]V, filterFunc func(K, V) bool) map[K]V {
	filtered := make(map[K]V)
	for k, v := range m {
		if filterFunc(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

func meetsCondition[T comparable](value T, condition FilterCondition, target T) (bool, error) {
	switch condition {
	case FilterEqualTo:
		return value == target, nil
	case FilterNotEqualTo:
		return value != target, nil
	case FilterGreaterThan, FilterLessThan, FilterGreaterThanOrEqualTo, FilterLessThanOrEqualTo:
		return compareAsNumeric(value, target, condition)
	case FilterContains, FilterNotContains:
		strValue, okValue := any(value).(string)
		strTarget, okTarget := any(target).(string)
		if !okValue || !okTarget {
			return false, fmt.Errorf("unsupported type for %v: %T", condition, value)
		}
		if condition == FilterContains {
			return containsString(strValue, strTarget), nil
		} else {
			return !containsString(strValue, strTarget), nil
		}
	default:
		return false, fmt.Errorf("unknown filter condition: %v", condition)
	}
}

// containsString 函數，用於檢查字串是否包含目標值
func containsString(value, target string) bool {
	return strings.Contains(value, target)
}

// compareAsNumeric 函數，用於處理 types.Numeric 類型的比較
func compareAsNumeric[T comparable](value T, target T, condition FilterCondition) (bool, error) {
	switch v := any(value).(type) {
	case int:
		return compareNumericValues(int64(v), int64(any(target).(int)), condition)
	case int8:
		return compareNumericValues(int64(v), int64(any(target).(int8)), condition)
	case int16:
		return compareNumericValues(int64(v), int64(any(target).(int16)), condition)
	case int32:
		return compareNumericValues(int64(v), int64(any(target).(int32)), condition)
	case int64:
		return compareNumericValues(v, any(target).(int64), condition)
	case uint:
		return compareNumericValues(uint64(v), uint64(any(target).(uint)), condition)
	case uint8:
		return compareNumericValues(uint64(v), uint64(any(target).(uint8)), condition)
	case uint16:
		return compareNumericValues(uint64(v), uint64(any(target).(uint16)), condition)
	case uint32:
		return compareNumericValues(uint64(v), uint64(any(target).(uint32)), condition)
	case uint64:
		return compareNumericValues(v, any(target).(uint64), condition)
	case float32:
		return compareNumericValues(float64(v), float64(any(target).(float32)), condition)
	case float64:
		return compareNumericValues(v, any(target).(float64), condition)
	default:
		return false, fmt.Errorf("unsupported type for %v: %T", condition, value)
	}
}

// compareNumericValues 函數，用於比較數字類型的值
func compareNumericValues[T types.Numeric](value T, target T, condition FilterCondition) (bool, error) {
	switch condition {
	case FilterGreaterThan:
		return value > target, nil
	case FilterLessThan:
		return value < target, nil
	case FilterGreaterThanOrEqualTo:
		return value >= target, nil
	case FilterLessThanOrEqualTo:
		return value <= target, nil
	default:
		return false, errors.New("invalid condition for numeric comparison")
	}
}
