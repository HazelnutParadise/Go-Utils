# Go-Utils

`Go-Utils` 是一個集合多個實用工具包的 Go 函式庫，旨在幫助開發者更高效地處理常見的開發任務。

## 工具包

### jsonutil

`jsonutil` 專門用於處理 JSON 文件。它提供了讀取 JSON 文件並解析為 `map[string]interface{}` 的功能，以及根據指定鍵路徑提取子 `map` 的功能。適合用於讀取 `config.json` 設定檔。

**功能：**

1. **LoadJSONFileToMap(filename string) (map[string]interface{}, error)**  
   將 JSON 文件加載為 `map[string]interface{}`。  
   - **參數：** `filename` - 要加載的 JSON 文件路徑。
   - **返回值：**
     - `map[string]interface{}`：解析後的字典結構。
     - `error`：如果加載或解析失敗，返回錯誤信息。

2. **LoadJSONFileAndExtractSubMap(filename string, path ...string) (map[string]interface{}, error)**  
   根據鍵路徑提取 JSON 文件中的子 `map`。  
   - **參數：** `filename` - 要加載的 JSON 文件路徑；`path` - 鍵路徑。
   - **返回值：**
     - `map[string]interface{}`：提取的子字典結構。
     - `error`：如果提取失敗，返回錯誤信息。

### mathutil

`mathutil` 提供了與數學運算相關的實用函數，例如適用於浮點數的四捨五入處理。

**功能：**

1. **RoundFloat64(value float64, precision int) float64**  
   將 `float64` 數值四捨五入到指定的小數位數。  
   - **參數：** `value` - 要四捨五入的浮點數；`precision` - 小數點後保留的位數。
   - **返回值：**
     - `float64`：四捨五入後的數值。

2. **RoundFloat32(value float32, precision int) float32**  
   將 `float32` 數值四捨五入到指定的小數位數。  
   - **參數：** `value` - 要四捨五入的浮點數；`precision` - 小數點後保留的位數。
   - **返回值：**
     - `float32`：四捨五入後的數值。

### sliceutil

`sliceutil` 專注於處理和操作切片數據結構。提供了針對數字切片和通用切片的多種實用函數。

**功能：**

1. **Max(slice []T) (T, error)**  
   返回數字切片中的最大值。  
   - **參數：** `slice` - 一個數字類型的切片。
   - **返回值：**
     - `T`：切片中的最大值。
     - `error`：如果切片為空，返回錯誤信息。

2. **Min(slice []T) (T, error)**  
   返回數字切片中的最小值。  
   - **參數：** `slice` - 一個數字類型的切片。
   - **返回值：**
     - `T`：切片中的最小值。
     - `error`：如果切片為空，返回錯誤信息。

3. **Average(slice []T) (float64, error)**  
   計算數字切片的算術平均值。  
   - **參數：** `slice` - 一個數字類型的切片。
   - **返回值：**
     - `float64`：切片的算術平均值。
     - `error`：如果切片為空，返回錯誤信息。

4. **Sort(slice []T, ascending ...bool) error**  
   對數字切片進行排序，根據 `ascending` 參數決定升序或降序，默認為升序。此函數會直接修改傳入的切片，而不返回新的副本。  
   - **參數：** `slice` - 一個數字類型的切片；`ascending` - 一個可選的布林值，默認為 `true`（升序）。
   - **返回值：**
     - `error`：如果傳入多個布林值，返回錯誤信息。

5. **Unique(slice []T) []T**  
   去除切片中的重複元素，適用於所有類型。  
   - **參數：** `slice` - 一個可比較類型的切片。
   - **返回值：**
     - `[]T`：去重後的切片。

6. **Reverse(slice []T)**  
   反轉切片中的元素順序。  
   - **參數：** `slice` - 一個任意類型的切片。
   - **返回值：** 無返回值，原地反轉。

7. **FindFirst(slice []T, target T) int**  
   查找第一個匹配的元素，返回其索引，未找到則返回 `-1`。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。
   - **返回值：**
     - `int`：首次匹配的索引，未找到返回 `-1`。

8. **FindAll(slice []T, target T) []int**  
   查找所有匹配的元素，返回其索引切片，未找到則返回空切片。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。
   - **返回值：**
     - `[]int`：所有匹配元素的索引切片。

9. **Contains(slice []T, target T) bool**  
   檢查切片中是否包含特定元素。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。
   - **返回值：**
     - `bool`：如果找到目標值，返回 `true`，否則返回 `false`。

10. **InsertAt(slice []T, index int, values ...T) ([]T, error)**  
    在指定位置插入元素，支持負索引。  
    - **參數：** `slice` - 一個任意類型的切片；`index` - 插入的位置，支持負索引；`values` - 要插入的值。
    - **返回值：**
      - `[]T`：插入後的切片。
      - `error`：如果索引無效，返回錯誤信息。

11. **Remove(slice []T, index int) ([]T, error)**  
    移除指定索引處的元素，支持負索引。  
    - **參數：** `slice` - 一個任意類型的切片；`index` - 要移除的位置，支持負索引。
    - **返回值：**
      - `[]T`：移除後的切片。
      - `error`：如果索引無效，返回錯誤信息。

12. **RemoveAll(slice []T, target T) []T**  
    移除切片中所有匹配目標的元素。  
    - **參數：** `slice` - 一個可比較類型的切片；`target` - 要移除的目標值。
    - **返回值：**
      - `[]T`：移除後的切片。

13. **Flatten(input interface{}) ([]T, error)**  
    將多層嵌套的切片展平成單層切片。  
    - **參數：** `input` - 可能包含多層嵌套的任意類型切片。
    - **返回值：**
      - `[]T`：展平後的單層切片。
      - `error`：如果遇到無法處理的類型，返回錯誤信息。

### maputil

`maputil` 專注於處理和操作 `map` 數據結構，提供了各種實用函數來簡化 `map` 的操作和管理。

**功能：**

1. **Keys(m map[K]V) []K**  
   返回 `map` 中的所有鍵。  
   - **參數：** `m` - 一個 `map`。
   - **返回值：**
     - `[]K`：包含所有鍵的切片。

2. **Values(m map[K]V) []V**  
   返回 `map` 中的所有值。  
   - **參數：** `m` - 一個 `map`。
   - **返回值：**
     - `[]V`：包含所有值的切片。

3. **Invert(m map[K]V) map[V]K**  
   反轉 `map` 的鍵和值的位置。  
   - **參數：** `m` - 一個 `map`。
   - **返回值：**
     - `map[V]K`：反轉後的新 `map`。

4. **FilterByKey(m map[K]V, condition FilterCondition, target K) (map[K]V, error)**  
   根據鍵來篩選 `map` 中的鍵值對。  
   - **參數：** `m` - 一個 `map`；`condition` - 篩選條件；`target` - 要比較的目標鍵。
   - **返回值：**
     - `map[K]V`：篩選後的 `map`。
     - `error`：如果篩選過程中出現錯誤，返回錯誤信息。

5. **FilterByValue(m map[K]V, condition FilterCondition, target V) (map[K]V, error)**  
   根據值來篩選 `map` 中的鍵值對。  
   - **參數：** `m` - 一個 `map`；`condition` - 篩選條件；`target` - 要比較的目標值。
   - **返回值：**
     - `map[K]V`：篩選後的 `map`。
     - `error`：如果篩選過程中出現錯誤，返回錯誤信息。

   **比較模式（FilterCondition）：**
   - `FilterEqualTo`：篩選等於目標值的鍵值對。
   - `FilterNotEqualTo`：篩選不等於目標值的鍵值對。
   - `FilterGreaterThan`：篩選大於目標值的鍵值對（僅適用於數字類型）。
   - `FilterLessThan`：篩選小於目標值的鍵值對（僅適用於數字類型）。
   - `FilterGreaterThanOrEqualTo`：篩選大於或等於目標值的鍵值對（僅適用於數字類型）。
   - `FilterLessThanOrEqualTo`：篩選小於或等於目標值的鍵值對（僅適用於數字類型）。
   - `FilterContains`：篩選包含目標值的字串鍵值對（僅適用於字串類型）。
   - `FilterNotContains`：篩選不包含目標值的字串鍵值對（僅適用於字串類型）。

6. **CustomFilter(m map[K]V, filterFunc func(K, V) bool) map[K]V**  
   使用自訂的篩選函數來篩選 `map` 中的鍵值對。  
   - **參數：** `m` - 一個 `map`；`filterFunc` - 自訂的篩選函數。
   - **返回值：**
     - `map[K]V`：篩選後的 `map`。

7. **Merge(m1, m2 map[K]V, opts ...interface{}) (map[K]V, error)**  
   合併兩個 `map`，在鍵衝突時根據指定的策略處理。當不傳入策略時，默認使用 `MergeDefault` 策略。  
   - **參數：** `m1` - 第一個 `map`；`m2` - 第二個 `map`；`opts` - 可選的合併策略或自訂的 resolver 函數。
   - **返回值：**
     - `map[K]V`：合併後的 `map`。
     - `error`：如果合併過程中出現錯誤，返回錯誤信息。

   **合併策略（MergeConflictResolutionStrategy）：**
   - `MergeDefault`：遇到衝突時返回錯誤。
   - `MergeUseFirst`：使用第一個 `map` 的值。
   - `MergeUseSecond`：使用第二個 `map` 的值。
   - `MergeAddValues`：將兩個值相加（僅適用於數字類型）。
   - `MergeCustomResolver`：使用自訂的 resolver 函數來處理衝突。

## 安裝

您可以使用以下命令來安裝 `Go-Utils`：

```bash
go get -u github.com/HazelnutParadise/Go-Utils
```

## 使用方法

在您的 Go 項目中導入 `Go-Utils`，然後使用所需的工具包來解決特定問題。

```go
import (
    "fmt"
    "log"
    "github.com/HazelnutParadise/Go-Utils/jsonutil"
    "github.com/HazelnutParadise/Go-Utils/mathutil"
    "github.com/HazelnutParadise/Go-Utils/sliceutil"
    "github.com/HazelnutParadise/Go-Utils/maputil"
)

func main() {
    data, err := jsonutil.LoadJSONFileToMap("config.json")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(data)

    rounded64 := mathutil.RoundFloat64(1.23456, 2)
    fmt.Println(rounded64) // 輸出: 1.23

    nums := []int{3, 1, 4, 1, 5, 9, 2}
    maxVal, _ := sliceutil.Max(nums)
    fmt.Println("Max:", maxVal) // 輸出: 9

    myMap := map[string]int{"a": 1, "b": 2, "c": 3}
    invertedMap := maputil.Invert(myMap)
    fmt.Println("Inverted Map:", invertedMap) // 輸出: map[1:a 2:b 3:c]
}
```

## 貢獻

我們歡迎您的貢獻！請通過提交 issue 或 pull request 來幫助我們改進此項目。如果有任何建議或問題，請隨時與我們聯繫。

## 許可證

此項目遵循 MIT 許可證 - 詳細內容請參見 [LICENSE](LICENSE) 文件。
