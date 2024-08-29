# Go-Utils

`Go-Utils` 是一個集合多個實用工具包的 Go 函式庫，旨在幫助開發者更高效地處理常見的開發任務。

## 設計原則

在 `Go-Utils` 中，所有操作數據結構的函數遵循以下設計原則：

1. **操作會改變數據值的函數會返回新數據結構。**
   - 例如：`Remove`、`InsertAt`、`Merge` 等函數會返回新的切片或 `map`。
  
2. **操作不會改變數據值的函數會直接修改原數據結構。**
   - 例如：`Sort`、`Reverse` 等函數直接對原數據結構進行操作，而不返回新的副本。

## 工具包

### conv

`conv` 包提供了一組簡單而強大的函數，用於將各種資料類型轉換為常用的 Go 基本類型。所有函數在轉換失敗時會直接 `panic`，確保操作的可靠性和一致性。

**功能：**

1. **ParseF64(value interface{}) float64**  
   將任意資料轉換為 `float64`，如果轉換失敗，則會 `panic`。  
   - **參數：** `value` - 任意可轉換為 `float64` 的資料。
   - **返回值：**
     - `float64`：轉換後的 `float64` 值。

2. **ParseF32(value interface{}) float32**  
   將任意資料轉換為 `float32`，如果轉換失敗，則會 `panic`。  
   - **參數：** `value` - 任意可轉換為 `float32` 的資料。
   - **返回值：**
     - `float32`：轉換後的 `float32` 值。

3. **ParseInt(value interface{}) int**  
   將任意資料轉換為 `int`，如果轉換失敗，則會 `panic`。  
   - **參數：** `value` - 任意可轉換為 `int` 的資料。
   - **返回值：**
     - `int`：轉換後的 `int` 值。

4. **ParseBool(value interface{}) bool**  
   將任意資料轉換為 `bool`，如果轉換失敗，則會 `panic`。  
   - **參數：** `value` - 任意可轉換為 `bool` 的資料。可以接受的值包括：
     - 字串 `"true"`, `"false"`, `"1"`, `"0"`, `"yes"`, `"no"`, `"on"`, `"off"`, `空字串` 等。
     - 任意數字類型，非零數字轉換為 `true`，零轉換為 `false`。
   - **返回值：**
     - `bool`：轉換後的 `bool` 值。

5. **ToString(value interface{}) string**  
   將任意資料轉換為字串，使用 `fmt.Sprintf` 進行格式化，錯誤時直接 `panic`。  
   - **參數：** `value` - 任意可轉換為字串的資料。
   - **返回值：**
     - `string`：轉換後的字串。

### errutil

`errutil` 包含一組處理錯誤的實用函數，旨在幫助開發者簡化錯誤處理流程，提高代碼的可讀性和可維護性。

**功能：**

1. **PanicOnErr(fn interface{}, args ...interface{}) []interface{}**  
   調用任意函數並自動處理返回的錯誤。如果該函數返回 `error`，且該 `error` 不為 `nil`，則 `PanicOnErr` 會觸發 `panic`，否則返回該函數的其他返回值。  
   - **參數：**  
     - `fn` - 需要調用的任意函數。
     - `args` - 傳遞給 `fn` 的參數列表。
   - **返回值：**  
     - `[]interface{}`：返回函數 `fn` 的所有非 `error` 返回值，包在一個切片裡。

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

3. **LoadJSONFileToStruct(filePath string, result interface{}) error**
   讀取 JSON 文件並將其解析為傳入的 Go 結構。適用於已經定義好的結構體。
   - **參數：** `filePath` - JSON 文件的路徑；`result` - 解析結果的 Go 結構體，需要提前定義好結構體來匹配 JSON 文件中的數據結構。
   - **返回值：**
     - `error` - 如果讀取或解析過程中出現錯誤，將返回錯誤信息。

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

3. **SplitFloat(value T, mode ...SplitFloatMode) (interface{}, interface{})**  
   根據指定模式將任何數字類型的變數（包括 `int`、`float32`、`float64` 等）分成整數部分和小數部分。  
   - **參數：**  
     - `value` - 任何數字類型的變數，泛型 `T` 可以是 `int`、`float32`、`float64` 等。
     - `mode` - 可選參數，指定返回結果的模式。若不指定模式，預設為 `SplitFloat_IntFloat`。傳入多個模式會觸發 `panic` 錯誤。
   - **返回值：**  
     - `interface{}`：返回兩個值，類型由選擇的模式決定：
       - **`SplitFloat_IntFloat`**: 返回 `int` 和 `float64`。
       - **`SplitFloat_IntInt`**: 返回 `int` 和 `int`，小數部分放大後取整。
       - **`SplitFloat_FloatFloat`**: 返回 `float64` 和 `float64`。

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

4. **Sum(slice []T) (T, error)**  
   計算數字切片中所有元素的總和。  
   - **參數：** `slice` - 一個數字類型的切片。
   - **返回值：**
     - `T`：切片中所有元素的總和。
     - `error`：如果切片為空，返回錯誤信息。

5. **Sort(slice []T, ascending ...bool) error**  
   對數字切片進行排序，根據 `ascending` 參數決定升序或降序，默認為升序。此函數會直接修改傳入的切片，而不返回新的副本。  
   - **參數：** `slice` - 一個數字類型的切片；`ascending` - 一個可選的布林值，默認為 `true`（升序）。
   - **返回值：**
     - `error`：如果傳入多個布林值，返回錯誤信息。

6. **Unique(slice []T) []T**  
   去除切片中的重複元素，適用於所有類型。  
   - **參數：** `slice` - 一個可比較類型的切片。
   - **返回值：**
     - `[]T`：去重後的切片。

7. **Reverse(slice []T)**  
   反轉切片中的元素順序。  
   - **參數：** `slice` - 一個任意類型的切片。
   - **返回值：** 無返回值，原地反轉。

8. **FindFirst(slice []T, target T) int**  
   查找第一個匹配的元素，返回其索引，未找到則返回 `-1`。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。
   - **返回值：**
     - `int`：首次匹配的索引，未找到返回 `-1`。

9. **FindAll(slice []T, target T) []int**  
   查找所有匹配的元素，返回其索引切片，未找到則返回空切片。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。
   - **返回值：**
     - `[]int`：所有匹配元素的索引切片。

10. **Contains(slice []T, target T) bool**  
   檢查切片中是否包含特定元素。  
   - **參數：** `slice` - 一個可比較類型的切片；`target` - 要查找的目標值。  
   - **返回值：**  
     - `bool`：如果找到目標值，返回 `true`，否則返回 `false`。  

11. **InsertAt(slice []T, index int, values ...T) ([]T, error)**  
    在指定位置插入元素，支持負索引。  
    - **參數：** `slice` - 一個任意類型的切片；`index` - 插入的位置，支持負索引；`values` - 要插入的值。
    - **返回值：**
      - `[]T`：插入後的切片。
      - `error`：如果索引無效，返回錯誤信息。

12. **Remove(slice []T, index int) ([]T, error)**  
    移除指定索引處的元素，支持負索引。  
    - **參數：** `slice` - 一個任意類型的切片；`index` - 要移除的位置，支持負索引。
    - **返回值：**
      - `[]T`：移除後的切片。
      - `error`：如果索引無效，返回錯誤信息。

13. **RemoveAll[T comparable](slice []T, targets ...T) []T**  
    移除切片中所有匹配目標的元素。  
    - **參數：** `slice` - 一個可比較類型的切片；`target` - 要移除的目標值（可以多個）。
    - **返回值：**
      - `[]T`：移除後的切片。

14. **Flatten(input interface{}) ([]T, error)**  
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

8. **RemoveKV(m map[K]V, key K, value V, ignoreErrors ...bool) (map[K]V, error)**  
   從 `map` 中移除指定的鍵值對並返回新的 `map`。  
   - **參數：** `m` - 一個 `map`；`key` - 要移除的鍵；`value` - 要移除的值；`ignoreErrors` - 可選，是否忽略錯誤。
   - **返回值：**
     - `map[K]V`：移除後的 `map`。
     - `error`：如果操作過程中出現錯誤，返回錯誤信息。

9. **RemoveByKey(m map[K]V, key K, ignoreErrors ...bool) (map[K]V, error)**  
   從 `map` 中移除指定的鍵並返回新的 `map`。  
   - **參數：** `m` - 一個 `map`；`key` - 要移除的鍵；`ignoreErrors` - 可選，是否忽略錯誤。
   - **返回值：**
     - `map[K]V`：移除後的 `map`。
     - `error`：如果操作過程中出現錯誤，返回錯誤信息。

10. **RemoveByValue(m map[K]V, value V, ignoreErrors ...bool) (map[K]V, error)**  
    從 `map` 中移除具有指定值的所有鍵值對並返回新的 `map`。  
    - **參數：** `m` - 一個 `map`；`value` - 要移除的值；`ignoreErrors` - 可選，是否忽略錯誤。
    - **返回值：**
      - `map[K]V`：移除後的 `map`。
      - `error`：如果操作過程中出現錯誤，返回錯誤信息。

11. **RemoveByMap(m map[K]V, toRemove map[K]V, ignoreErrors ...bool) (map[K]V, error)**  
    根據傳入的鍵值對 `map` 刪除 `map` 中的對應鍵值對。  
    - **參數：** `m` - 一個 `map`；`toRemove` - 包含要移除的鍵值對的 `map`；`ignoreErrors` - 可選，是否忽略錯誤。
    - **返回值：**
      - `map[K]V`：移除後的 `map`。
      - `error`：如果操作過程中出現錯誤，返回錯誤信息。

12. **RemoveByKeys(m map[K]V, keys []K, ignoreErrors ...bool) (map[K]V, error)**  
    從 `map` 中移除所有指定鍵並返回新的 `map`。  
    - **參數：** `m` - 一個 `map`；`keys` - 包含要移除的鍵的切片；`ignoreErrors` - 可選，是否忽略錯誤。
    - **返回值：**
      - `map[K]V`：移除後的 `map`。
      - `error`：如果操作過程中出現錯誤，返回錯誤信息。

13. **RemoveByValues(m map[K]V, values []V, ignoreErrors ...bool) (map[K]V, error)**  
    從 `map` 中移除所有具有指定值的鍵值對並返回新的 `map`。  
    - **參數：** `m` - 一個 `map`；`values` - 包含要移除的值的切片；`ignoreErrors` - 可選，是否忽略錯誤。
    - **返回值：**
      - `map[K]V`：移除後的 `map`。
      - `error`：如果操作過程中出現錯誤，返回錯誤信息。

### timeutil

`timeutil` 提供了與時間相關的實用函數，涵蓋時間格式化、時區轉換以及日期計算等常見操作。

**功能：**

1. **TimeInZone(offsetHours int) time.Time**  
   根據指定的 UTC 偏移值返回該時區的當前時間。  
   - **參數：** `offsetHours` - UTC 的偏移值，以小時為單位，正數代表東部時區，負數代表西部時區。
   - **返回值：**
     - `time.Time`：該時區的當前時間。

2. **NowFormatted(format string, timezoneOffset ...int) string**  
   根據指定的格式和時區偏移量返回當前時間的字串表示，預設使用 UTC+0。如果提供了多個時區偏移值，將會 `panic`。  
   - **參數：**  
     - `format` - 時間格式字串，例如 `"2006-01-02 15:04:05"`。  
     - `timezoneOffset` - 可選參數，UTC 的偏移值，以小時為單位，預設為 UTC+0。
   - **返回值：**
     - `string`：格式化後的當前時間字串。

3. **FormatTime(t time.Time, format string) string**  
   將指定的時間根據指定的格式進行格式化。  
   - **參數：**  
     - `t` - 要格式化的時間。  
     - `format` - 時間格式字串，例如 `"2006-01-02 15:04:05"`。
   - **返回值：**
     - `string`：格式化後的時間字串。

4. **DaysBetween(startDate, endDate time.Time) int**  
   計算兩個日期之間的天數，返回正數。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的天數。

5. **DaysDiff(startDate, endDate time.Time) int**  
   計算兩個日期之間的天數，不取絕對值。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的天數。

6. **MonthsBetween(startDate, endDate time.Time) int**  
   計算兩個日期之間的月份數，返回正數。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的月份數。

7. **MonthsDiff(startDate, endDate time.Time) int**  
   計算兩個日期之間的月份數，不取絕對值。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的月份數。

8. **YearsBetween(startDate, endDate time.Time) int**  
   計算兩個日期之間的年數，返回正數。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的年數。

9. **YearsDiff(startDate, endDate time.Time) int**  
   計算兩個日期之間的年數，不取絕對值。  
   - **參數：**  
     - `startDate` - 起始日期。  
     - `endDate` - 結束日期。
   - **返回值：**
     - `int`：兩個日期之間的年數。

**常用時間格式（可代替格式字串）：**

- `FormatDateOnly`: `"2006-01-02"` - 只顯示日期。
- `FormatTimeOnly`: `"15:04:05"` - 只顯示時間。
- `FormatDateTime`: `"2006-01-02 15:04:05"` - 顯示日期和時間。
- `FormatISO8601`: `"2006-01-02T15:04:05Z07:00"` - ISO8601 格式。
- `FormatISO8601Compact`: `"20060102T150405Z0700"` - 緊湊版的 ISO8601 格式。
- `FormatRFC1123`: RFC1123 格式。
- `FormatRFC822`: RFC822 格式。

10. **UnixAfterSeconds(seconds int) int64**  
   根據當前時間計算指定秒數後的 Unix 時間戳，如果秒數為負，則計算指定秒數前的 Unix 時間戳。
   - **參數：** 
     - `seconds` - 指定的秒數，可以是正、零或負數。
   - **返回值：** 
     - `int64` - 指定秒數後或前的 Unix 時間戳。 

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
