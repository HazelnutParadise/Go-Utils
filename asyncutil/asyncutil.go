package asyncutil

import (
	"reflect"
	"sync"
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

// ParallelProcess 接受一個 map，其中鍵是要平行處理的函數，值是該函數的參數切片。
// 函數會平行執行所有的函數並返回結果。
func ParallelProcess(tasks map[interface{}][]interface{}) map[interface{}][]interface{} {
	results := make(map[interface{}][]interface{})
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for fn, args := range tasks {
		wg.Add(1)
		go func(fn interface{}, args []interface{}) {
			defer wg.Done()
			fnValue := reflect.ValueOf(fn)
			in := make([]reflect.Value, len(args))
			for i, arg := range args {
				in[i] = reflect.ValueOf(arg)
			}

			out := fnValue.Call(in)

			// 轉換結果為 interface{} 切片
			result := make([]interface{}, len(out))
			for i, val := range out {
				result[i] = val.Interface()
			}

			// 使用鎖來保護 results map
			mu.Lock()
			results[fn] = result
			mu.Unlock()
		}(fn, args)
	}

	wg.Wait()
	return results
}
