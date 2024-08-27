// sliceutil/generic.go
package sliceutil

import (
	"errors"
)

// Unique 函數，去除切片中的重複元素，適用於所有類型
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := []T{}

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// Reverse 函數，反轉切片中元素的順序
func Reverse[T any](slice []T) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
}

// FindFirst 函數，查找第一個匹配的元素，返回索引值，未找到則返回 -1
func FindFirst[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1 // 未找到返回 -1
}

// FindAll 函數，查找所有匹配的元素，返回索引切片，未找到則返回空切片
func FindAll[T comparable](slice []T, target T) []int {
	var indices []int
	for i, v := range slice {
		if v == target {
			indices = append(indices, i)
		}
	}
	return indices // 未找到則返回空切片
}

// Contains 函數，檢查切片中是否包含特定元素
func Contains[T comparable](slice []T, target T) bool {
	found := FindFirst(slice, target)
	return found != -1
}

// InsertAt 函數，允許在指定位置插入元素，支持負索引
// 插入後，指定索引處的元素是新插入的元素，負索引也適用
func InsertAt[T any](slice []T, index int, values ...T) ([]T, error) {
	if index < 0 {
		index = len(slice) + index + 1
	}
	if index < 0 || index > len(slice) {
		return slice, errors.New("index out of bounds")
	}
	// 插入元素，保持指定索引不變
	return append(slice[:index], append(values, slice[index:]...)...), nil
}

// Remove 函數，移除指定索引處的元素，若索引超出範圍則返回錯誤，支持負索引
func Remove[T any](slice []T, index int) ([]T, error) {
	if index < 0 {
		index = len(slice) + index
	}
	if index < 0 || index >= len(slice) {
		return slice, errors.New("index out of bounds")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

// RemoveAll 函數，移除切片中所有等於 target 的元素
func RemoveAll[T comparable](slice []T, target T) []T {
	result := slice[:0] // 保留原切片的容量
	for _, v := range slice {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}

// Flatten 函數，將多層嵌套的切片展平成單層切片
func Flatten[T any](input interface{}) ([]T, error) {
	var result []T

	switch v := input.(type) {
	case []interface{}:
		for _, item := range v {
			flattened, err := Flatten[T](item)
			if err != nil {
				return nil, err
			}
			result = append(result, flattened...)
		}
	case T:
		result = append(result, v)
	default:
		return nil, errors.New("unexpected type encountered")
	}

	return result, nil
}
