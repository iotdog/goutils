package auth

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 哈希密码字符串，成功则返回十六进制字符串，否则返回空字符串
func HashPassword(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

// CheckPasswordHash 检查密码是否与hash值一致
func CheckPasswordHash(pwd, hash string) bool {
	// 先将十六进制hash字符串转换为bytes
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(pwd))
	return err == nil
}
