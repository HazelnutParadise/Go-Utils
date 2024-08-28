package errutil

import "reflect"

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
