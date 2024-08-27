// types/numeric.go
package types

// Numeric 定義僅適用於數字類型的泛型約束
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
