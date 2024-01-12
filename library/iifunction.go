package library

import "golang.org/x/crypto/bcrypt"

// 密码加密
func HashString(hashString string) (string, error) {
	// GenerateFromPassword 返回密码的bcrypt哈希值
	hash, err := bcrypt.GenerateFromPassword([]byte(hashString), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
