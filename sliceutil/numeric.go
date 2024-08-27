// sliceutil/numeric.go
package sliceutil

import (
	"errors"
	"sort"
)

// 定義一個泛型約束，涵蓋所有數字類型
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Max 函數，返回切片中的最大值
func Max[T Numeric](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("slice is empty")
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Min 函數，返回切片中的最小值
func Min[T Numeric](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("slice is empty")
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// Average 函數，計算算術平均值
func Average[T Numeric](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice)), nil
}

// Sort 函數，對數字切片排序，根據 ascending 參數決定升序或降序，默認為升序
func Sort[T Numeric](slice []T, ascending ...bool) error {
	if len(ascending) > 1 {
		return errors.New("the Sort function allows only one boolean value for the ascending parameter")
	}

	asc := true // 默認值為 true，即升序
	if len(ascending) == 1 {
		asc = ascending[0]
	}

	if asc {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	} else {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})
	}
	return nil
}
