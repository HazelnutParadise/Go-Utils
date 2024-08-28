package mathutil

import (
	"math"

	"github.com/HazelnutParadise/Go-Utils/types" // 假設有一個定義 Numeric 類型約束的包
)

// SplitFloatMode 定義枚舉類型，用於指定 SplitFloat 的返回模式
type SplitFloatMode int

const (
	SplitFloat_IntFloat   SplitFloatMode = iota // 返回 int 和 float64
	SplitFloat_IntInt                           // 返回 int 和 int
	SplitFloat_FloatFloat                       // 返回 float64 和 float64
)

// SplitFloat 函數，根據指定模式將浮點數分成整數部分和小數部分
func SplitFloat[T types.Numeric](value T, mode ...SplitFloatMode) (interface{}, interface{}) {
	selectedMode := SplitFloat_IntFloat // 設置預設模式為 SplitFloat_IntFloat

	// 如果傳入了模式參數，則使用該參數
	if len(mode) == 1 {
		selectedMode = mode[0]
	} else if len(mode) > 1 {
		panic("SplitFloat: too many arguments, only one mode can be specified")
	}

	// 將 value 轉換為 float64 以進行數學計算
	floatValue := float64(value)
	intPart := math.Floor(floatValue)
	fracPart := floatValue - intPart

	switch selectedMode {
	case SplitFloat_IntInt:
		return int(intPart), int(fracPart * 1000) // 將小數部分放大後取整
	case SplitFloat_FloatFloat:
		return intPart, fracPart
	case SplitFloat_IntFloat:
		fallthrough
	default:
		return int(intPart), fracPart
	}
}
