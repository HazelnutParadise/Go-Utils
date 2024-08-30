package asyncutil

import (
	"reflect"
)

// Awaitable 表示一個可以等待的結果
type Awaitable struct {
	results []interface{}
	err     error
	done    chan struct{}
}

// NewAwaitable 創建一個新的 Awaitable
func NewAwaitable(fn interface{}, args ...interface{}) *Awaitable {
	a := &Awaitable{
		done: make(chan struct{}),
	}

	go func() {
		fnValue := reflect.ValueOf(fn)
		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}

		out := fnValue.Call(in)

		a.results = make([]interface{}, 0, len(out))
		for _, val := range out {
			// 檢查是否是 error 類型
			if err, ok := val.Interface().(error); ok {
				a.err = err
			} else {
				a.results = append(a.results, val.Interface())
			}
		}

		close(a.done)
	}()

	return a
}

// Await 等待結果，返回結果切片和 error
func (a *Awaitable) Await() ([]interface{}, error) {
	<-a.done
	return a.results, a.err
}

// Async 創建一個異步操作，並返回 Awaitable
func Async(fn interface{}, args ...interface{}) *Awaitable {
	return NewAwaitable(fn, args...)
}
