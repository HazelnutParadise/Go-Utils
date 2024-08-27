// maputil/merge.go
package maputil

import (
	"errors"
	"fmt"
)

// MergeConflictResolutionStrategy 定義鍵衝突時的處理策略
type MergeConflictResolutionStrategy int

const (
	MergeDefault        MergeConflictResolutionStrategy = iota // 默認策略：遇到衝突時返回錯誤
	MergeUseFirst                                              // 使用第一個 map 的值
	MergeUseSecond                                             // 使用第二個 map 的值
	MergeAddValues                                             // 將兩個值相加（僅適用於數字類型）
	MergeCustomResolver                                        // 使用自訂的 resolver 函數
)

// Merge 函數，合併兩個 map，在鍵衝突時根據指定的策略處理
// 當不傳入策略時，默認使用 MergeDefault 策略
func Merge[K comparable, V any](m1, m2 map[K]V, opts ...interface{}) (map[K]V, error) {
	merged := make(map[K]V, len(m1)+len(m2))

	// 設定默認策略為 MergeDefault
	strategy := MergeDefault
	var resolver func(V, V) V

	// 判斷是否傳入了策略和 resolver
	for _, opt := range opts {
		switch v := opt.(type) {
		case MergeConflictResolutionStrategy:
			strategy = v
		case func(V, V) V:
			resolver = v
		}
	}

	// 檢查當使用 MergeCustomResolver 策略時，是否提供了 resolver
	if strategy == MergeCustomResolver && resolver == nil {
		return nil, errors.New("resolver function must be provided for MergeCustomResolver strategy")
	}

	// 將 m1 的內容複製到 merged
	for k, v := range m1 {
		merged[k] = v
	}

	// 處理 m2 的內容
	for k, v := range m2 {
		if existing, ok := merged[k]; ok {
			switch strategy {
			case MergeDefault:
				return nil, fmt.Errorf("conflict detected on key: %v", k)
			case MergeUseFirst:
				// 保持原來的值，不做變動
			case MergeUseSecond:
				merged[k] = v // 使用 m2 的值覆蓋
			case MergeAddValues:
				switch any(existing).(type) {
				case int:
					merged[k] = any(any(existing).(int) + any(v).(int)).(V)
				case float64:
					merged[k] = any(any(existing).(float64) + any(v).(float64)).(V)
				default:
					return nil, errors.New("MergeAddValues strategy only supports int or float64 types")
				}
			case MergeCustomResolver:
				merged[k] = resolver(existing, v)
			default:
				return nil, errors.New("unknown conflict resolution strategy")
			}
		} else {
			merged[k] = v
		}
	}

	return merged, nil
}
