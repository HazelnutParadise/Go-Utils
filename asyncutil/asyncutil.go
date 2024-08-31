package asyncutil

import (
	"reflect"
	"runtime"
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

// getDefaultGoroutines 取得預設的線程數
func getDefaultGoroutines() int {
	numCPU := runtime.NumCPU()
	if numCPU > 0 {
		return numCPU
	}
	return 1 // 如果無法取得 CPU 核心數量，預設使用 1 個線程
}

// ParallelFor 用於平行處理 for 迴圈，支援切片和 map
func ParallelFor[T any](data interface{}, task func(T) interface{}, numGoroutines ...int) []interface{} {
	value := reflect.ValueOf(data)
	kind := value.Kind()

	// 確認是否是支援的類型
	if kind != reflect.Slice && kind != reflect.Map {
		panic("ParallelFor: unsupported data type, must be slice or map")
	}

	// 檢查是否有多個線程數參數
	if len(numGoroutines) > 1 {
		panic("ParallelFor: only one goroutine count can be specified")
	}

	// 取得預設的線程數
	goroutines := getDefaultGoroutines()
	if len(numGoroutines) == 1 {
		goroutines = numGoroutines[0]
	}

	length := value.Len()
	results := make([]interface{}, length)

	// 決定每個線程處理的數據量
	chunkSize := length / goroutines
	if length%goroutines != 0 {
		chunkSize++
	}

	var wg sync.WaitGroup

	if kind == reflect.Slice {
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				begin := i * chunkSize
				finish := begin + chunkSize
				if finish > length {
					finish = length
				}
				for j := begin; j < finish; j++ {
					results[j] = task(value.Index(j).Interface().(T))
				}
			}(i)
		}
	} else if kind == reflect.Map {
		keys := value.MapKeys()
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				begin := i * chunkSize
				finish := begin + chunkSize
				if finish > length {
					finish = length
				}
				for j := begin; j < finish; j++ {
					k := keys[j].Interface().(T)
					results[j] = task(k)
				}
			}(i)
		}
	}

	wg.Wait()

	return results
}

// ParallelForEach 用於平行處理 for range 迴圈，支援切片和 map
func ParallelForEach(data interface{}, task interface{}, numGoroutines ...int) []interface{} {
	dataValue := reflect.ValueOf(data)
	taskValue := reflect.ValueOf(task)

	dataKind := dataValue.Kind()
	if dataKind != reflect.Slice && dataKind != reflect.Map {
		panic("ParallelForEach: data must be a slice or map")
	}

	if len(numGoroutines) > 1 {
		panic("ParallelForEach: only one goroutine count can be specified")
	}

	goroutines := getDefaultGoroutines()
	if len(numGoroutines) == 1 {
		goroutines = numGoroutines[0]
	}

	length := dataValue.Len()
	results := make([]interface{}, length)

	chunkSize := length / goroutines
	if length%goroutines != 0 {
		chunkSize++
	}

	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin := i * chunkSize
			finish := begin + chunkSize
			if finish > length {
				finish = length
			}
			for j := begin; j < finish; j++ {
				var result reflect.Value
				if dataKind == reflect.Slice {
					result = taskValue.Call([]reflect.Value{reflect.ValueOf(j), dataValue.Index(j)})[0]
				} else if dataKind == reflect.Map {
					key := dataValue.MapKeys()[j]
					result = taskValue.Call([]reflect.Value{key, dataValue.MapIndex(key)})[0]
				}
				results[j] = result.Interface()
			}
		}(i)
	}
	wg.Wait()

	return results
}
