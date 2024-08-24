# Go-Utils

`Go-Utils` 是一個集合多個實用工具包的 Go 庫，旨在幫助開發者更高效地處理常見的開發任務。

## 工具包

### jsonutil

`jsonutil` 是 `Go-Utils` 中的一個工具包，專門用於處理 JSON 文件。它提供了讀取 JSON 文件並解析為 `map[string]interface{}` 的功能，以及根據指定鍵路徑提取子 `map` 的功能。適合用於讀取 `config.json` 設定檔。

**功能：**

1. **LoadJSONFileToMap** - 將 JSON 文件加載為 `map[string]interface{}`。
2. **LoadJSONFileAndExtractSubMap** - 根據鍵路徑提取 JSON 文件中的子 `map`。

**範例使用：**

```go
data, err := jsonutil.LoadJSONFileToMap("config.json")
if err != nil {
    log.Fatal(err)
}
fmt.Println(data)

serverConfig, err := jsonutil.LoadJSONFileAndExtractSubMap("config.json", "server")
if err != nil {
    log.Fatal(err)
}
fmt.Println(serverConfig)
```

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
)

func main() {
    data, err := jsonutil.LoadJSONFileToMap("config.json")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(data)
}
```

## 其他工具包

- **fileutil** - 處理文件讀寫操作。
- **netutil** - 簡化網絡請求和響應的操作。
- **dbutil** - 提供數據庫操作的便捷方法。

## 貢獻

我們歡迎您的貢獻！請通過提交 issue 或 pull request 來幫助我們改進此項目。如果有任何建議或問題，請隨時與我們聯繫。

## 許可證

此項目遵循 MIT 許可證 - 詳細內容請參見 [LICENSE](LICENSE) 文件。
