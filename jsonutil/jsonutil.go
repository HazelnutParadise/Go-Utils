package jsonutil

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// LoadJSONFileToMap 讀取 JSON 文件並將其解析為 map[string]interface{}
func LoadJSONFileToMap(filePath string) (map[string]interface{}, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// LoadJSONFileAndExtractSubMap 讀取 JSON 文件，並根據鍵路徑返回對應的子 map
func LoadJSONFileAndExtractSubMap(filePath string, keys ...string) (map[string]interface{}, error) {
	data, err := LoadJSONFileToMap(filePath)
	if err != nil {
		return nil, err
	}

	currentMap := data
	for _, key := range keys {
		if val, exists := currentMap[key]; exists {
			switch v := val.(type) {
			case map[string]interface{}:
				currentMap = v
			default:
				return nil, errors.New("the key '" + key + "' does not point to a map")
			}
		} else {
			return nil, errors.New("the key '" + key + "' was not found in the JSON structure")
		}
	}
	return currentMap, nil
}

func LoadJSONFileToStruct(filePath string, result interface{}) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, result); err != nil {
		return err
	}

	return nil
}
