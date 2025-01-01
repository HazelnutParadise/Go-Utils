// Description: 用於將文字雜湊化，並比較文字與雜湊後的文字是否相符
package hashutil

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

// HashConfig 用於設定雜湊的參數
type HashConfig struct {
	Memory     uint32 // 記憶體使用量，以 KB 為單位，預設值為 64 * 1024
	Iterations uint32 // 迭代次數，預設值為 3
	Threads    uint8  // 使用的執行緒數，預設值為 4
	CustomSalt []byte // 自訂鹽值，若未指定，則會自動生成
	SaltLength uint32 // 鹽值長度，當 CustomSalt 不為空時無效，預設值為 16
	KeyLength  uint32 // 雜湊後的字串長度，預設值為 32
}

var defaultHashConfig = HashConfig{
	Memory:     64 * 1024,
	Iterations: 3,
	Threads:    4,
	SaltLength: 16,
	KeyLength:  32,
}

// Hash 用於將文字雜湊化，您可以傳入 HashConfig 結構體來設定雜湊的參數
func Hash(text string, config ...HashConfig) (hashedBase64 string, saltBase64 string, err error) {
	if len(config) > 1 {
		return "", "", errors.New("only one config is allowed")
	}
	var c HashConfig
	if len(config) == 0 {
		c = defaultHashConfig
	} else {
		c = config[0]
	}

	var salt []byte
	if c.CustomSalt != nil {
		salt = c.CustomSalt
	} else {
		salt, err = generateSalt(c.SaltLength)
		if err != nil {
			return "", "", err
		}
	}
	hashed := argon2.IDKey([]byte(text), salt, c.Iterations, c.Memory, c.Threads, c.KeyLength)
	hashedBase64 = base64.RawStdEncoding.EncodeToString(hashed)
	saltBase64 = base64.RawStdEncoding.EncodeToString(salt)
	return hashedBase64, saltBase64, nil
}

func generateSalt(saltLength uint32) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// CompareHash 用於比較文字與雜湊後的文字是否相符，您可以傳入 HashConfig 結構體來設定雜湊的參數
func CompareHash(textToCompare, hashedBase64, saltBase64 string, hashConfig ...HashConfig) bool {
	if len(hashConfig) > 1 {
		return false
	}
	var c HashConfig
	if len(hashConfig) == 0 {
		c = defaultHashConfig
	} else {
		c = hashConfig[0]
	}

	saltBytes, err := base64.RawStdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false
	}
	hashedBytes, err := base64.RawStdEncoding.DecodeString(hashedBase64)
	if err != nil {
		return false
	}
	newHash := argon2.IDKey([]byte(textToCompare), saltBytes, c.Iterations, c.Memory, c.Threads, c.KeyLength)
	return bytes.Equal(newHash, hashedBytes)
}
