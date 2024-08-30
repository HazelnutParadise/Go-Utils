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

// Task 結構體，包含要執行的函數、其對應的參數和標識符
type Task struct {
	ID   string        // 任務的標識符
	Fn   interface{}   // 要執行的函數
	Args []interface{} // 函數的參數切片
}

// TaskResult 結構體，包含每個任務的結果和標識符
type TaskResult struct {
	ID      string        // 任務的標識符
	Results []interface{} // 函數返回的結果
}

// ParallelProcess 接受一個 Task 切片，平行執行所有的函數並返回結果。
func ParallelProcess(tasks []Task) []TaskResult {
	results := make([]TaskResult, len(tasks))
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task Task) {
			defer wg.Done()
			fnValue := reflect.ValueOf(task.Fn)
			in := make([]reflect.Value, len(task.Args))
			for j, arg := range task.Args {
				in[j] = reflect.ValueOf(arg)
			}

			out := fnValue.Call(in)

			// 轉換結果為 interface{} 切片
			result := make([]interface{}, len(out))
			for j, val := range out {
				result[j] = val.Interface()
			}

			results[i] = TaskResult{
				ID:      task.ID,
				Results: result,
			}
		}(i, task)
	}

	wg.Wait()
	return results
}
