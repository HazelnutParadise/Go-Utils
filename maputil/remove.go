// maputil/remove.go
package maputil

import (
	"errors"
	"fmt"
)

// RemoveKV 函數，從 map 中移除指定的鍵值對並返回新的 map
// 如果鍵不存在或者值不匹配，則報錯。可以選擇忽略錯誤。
func RemoveKV[K comparable, V comparable](m map[K]V, key K, value V, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	// 創建一個新的 map 來存儲結果
	newMap := make(map[K]V, len(m))
	for k, v := range m {
		if k != key || v != value {
			newMap[k] = v
		}
	}

	// 如果鍵不存在或者值不匹配，返回錯誤（除非忽略錯誤模式啟用）
	if originalValue, exists := m[key]; !exists || originalValue != value {
		if !ignore {
			return newMap, fmt.Errorf("key-value pair not found in map: %v=%v", key, value)
		}
	}

	return newMap, nil
}

// RemoveByKey 函數，從 map 中移除指定的鍵並返回新的 map
// 如果鍵不存在，則報錯。可以選擇忽略錯誤。
func RemoveByKey[K comparable, V any](m map[K]V, key K, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	// 創建一個新的 map 來存儲結果
	newMap := make(map[K]V, len(m))
	for k, v := range m {
		if k != key {
			newMap[k] = v
		}
	}

	// 如果鍵不存在，返回錯誤（除非忽略錯誤模式啟用）
	if _, exists := m[key]; !exists {
		if !ignore {
			return newMap, fmt.Errorf("key not found in map: %v", key)
		}
	}

	return newMap, nil
}

// RemoveByValue 函數，從 map 中移除具有指定值的所有鍵值對並返回新的 map
// 如果沒有找到該值，則報錯。可以選擇忽略錯誤。
func RemoveByValue[K comparable, V comparable](m map[K]V, value V, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	newMap := make(map[K]V, len(m))
	count := 0

	for k, v := range m {
		if v != value {
			newMap[k] = v
		} else {
			count++
		}
	}

	// 如果沒有找到該值，返回錯誤（除非忽略錯誤模式啟用）
	if count == 0 && !ignore {
		return newMap, fmt.Errorf("value not found in map: %v", value)
	}

	return newMap, nil
}

// RemoveByMap 函數，根據傳入的鍵值對 map 刪除另一個 map 中的對應鍵值對
// 默認情況下，當條件 map 中有任何鍵在目標 map 中不存在時會報錯。
// 如果設置 ignoreErrors 為 true，則忽略錯誤，只刪除匹配的部分。
func RemoveByMap[K comparable, V comparable](m map[K]V, toRemove map[K]V, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	// 創建一個新的 map 來存儲結果
	newMap := make(map[K]V, len(m))

	for k, v := range toRemove {
		if originalValue, exists := m[k]; !exists || originalValue != v {
			if !ignore {
				return nil, fmt.Errorf("key %v with value %v not found in map", k, v)
			}
		} else {
			delete(m, k)
		}
	}

	// 將剩餘的鍵值對複製到新 map
	for k, v := range m {
		newMap[k] = v
	}

	return newMap, nil
}

// RemoveByKeys 函數，從 map 中移除所有指定鍵並返回新的 map
// 如果有任何鍵不存在，則報錯。可以選擇忽略錯誤。
func RemoveByKeys[K comparable, V any](m map[K]V, keys []K, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	newMap := make(map[K]V, len(m))
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			if !ignore {
				return nil, fmt.Errorf("key not found in map: %v", key)
			}
		} else {
			delete(m, key)
		}
	}

	// 將剩餘的鍵值對複製到新 map
	for k, v := range m {
		newMap[k] = v
	}

	return newMap, nil
}

// RemoveByValues 函數，從 map 中移除所有具有指定值的鍵值對並返回新的 map
// 如果沒有找到任何指定值，則報錯。可以選擇忽略錯誤。
func RemoveByValues[K comparable, V comparable](m map[K]V, values []V, ignoreErrors ...bool) (map[K]V, error) {
	// 檢查是否有多於一個 ignoreErrors 參數
	if len(ignoreErrors) > 1 {
		return nil, errors.New("only one boolean value is allowed for ignoreErrors parameter")
	}

	// 默認情況下不忽略錯誤
	ignore := false
	if len(ignoreErrors) == 1 {
		ignore = ignoreErrors[0]
	}

	newMap := make(map[K]V, len(m))
	count := 0

	for _, value := range values {
		for k, v := range m {
			if v == value {
				delete(m, k)
				count++
			}
		}
	}

	// 如果沒有找到任何指定值，返回錯誤（除非忽略錯誤模式啟用）
	if count == 0 && !ignore {
		return newMap, fmt.Errorf("none of the values found in map: %v", values)
	}

	// 將剩餘的鍵值對複製到新 map
	for k, v := range m {
		newMap[k] = v
	}

	return newMap, nil
}
