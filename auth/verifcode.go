package auth

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateVerifCode 生成验证码字符串，包含六位随机数字
func GenerateVerifCode() string {
	rdsrc := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%d", 100000+int(900000*rdsrc.Float64()))
}
