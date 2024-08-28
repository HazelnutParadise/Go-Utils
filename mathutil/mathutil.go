package mathutil

import (
	"strings"

	"github.com/HazelnutParadise/Go-Utils/conv"
	"github.com/HazelnutParadise/Go-Utils/types"
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
	// 使用 ToString 函數將 float64 的小數部分轉換為字串
	floatValueStr := conv.ToString(floatValue)
	// 使用 Split 函數將整數部分和小數部分分開
	strs := strings.Split(floatValueStr, ".")
	intPartStr := strs[0]
	fracPartStr := ""
	if len(strs) == 2 {
		fracPartStr = strs[1]
	} else {
		fracPartStr = "0"
	}
	// 去掉小數部分尾隨的0
	fracPartStr = strings.TrimRight(fracPartStr, "0")
	if fracPartStr == "" {
		fracPartStr = "0"
	}

	switch selectedMode {
	case SplitFloat_IntInt:

		// 使用 ParseInt 函數將整數部分和小數部分轉換為整數
		intPart := conv.ParseInt(intPartStr)
		fracPart := conv.ParseInt(fracPartStr)
		return intPart, fracPart
	case SplitFloat_FloatFloat:
		// 使用 ParseFloat 函數將整數部分和小數部分轉換為浮點數
		intPart := conv.ParseF64(intPartStr)
		fracPartStr = "0." + fracPartStr
		fracPart := conv.ParseF64(fracPartStr)
		return intPart, fracPart
	case SplitFloat_IntFloat:
		fallthrough
	default:
		// 使用 ParseInt 函數將整數部分轉換為整數
		intPart := conv.ParseInt(intPartStr)
		// 使用 ParseF64 函數將小數部分轉換為浮點數
		fracPartStr = "0." + fracPartStr
		fracPart := conv.ParseF64(fracPartStr)
		return intPart, fracPart
	}
}
