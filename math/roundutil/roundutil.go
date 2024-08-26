package roundutil

import (
	"math"
)

// RoundFloat64 將 float64 數值四捨五入到指定的小數位數
// number: 需要四捨五入的浮點數
// precision: 保留的小數位數
func RoundFloat64(number float64, precision int) float64 {
	// 計算放大因子，例如 precision 為 2 時，factor = 100
	factor := math.Pow(10, float64(precision))
	// 將數值放大後進行四捨五入，然後再縮小回原來的精度
	return math.Round(number*factor) / factor
}

// RoundFloat32 將 float32 數值四捨五入到指定的小數位數
// number: 需要四捨五入的浮點數
// precision: 保留的小數位數
func RoundFloat32(number float32, precision int) float32 {
	// 計算放大因子，例如 precision 為 2 時，factor = 100
	factor := float32(math.Pow(10, float64(precision)))
	// 將數值放大後進行四捨五入，然後再縮小回原來的精度
	return float32(math.Round(float64(number*factor))) / factor
}
