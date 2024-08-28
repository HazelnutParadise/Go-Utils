package errutil

import "reflect"

// PanicOnErr 用於處理函數調用時的錯誤，如果函數返回的錯誤不為 nil，則直接 panic
// fn: 要調用的函數
// args: 要傳遞給函數的參數
// 返回值: 函數的返回值切片
// 需自行處理函數返回值的類型
func PanicOnErr(fn interface{}, args ...interface{}) []interface{} {
	// 將傳入的函數轉換為反射對象
	fnValue := reflect.ValueOf(fn)

	// 構建參數列表
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// 調用函數並獲取返回值
	out := fnValue.Call(in)

	// 遍歷所有返回值
	var results []interface{}
	for _, val := range out {
		if val.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if !val.IsNil() {
				panic(val.Interface().(error))
			}
		} else {
			results = append(results, val.Interface())
		}
	}

	return results
}
