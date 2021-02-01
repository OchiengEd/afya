package platform

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,-_+=><*|"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PasswordGen generates passwords
func PasswordGen(length int) string {
	passwd := make([]byte, length)
	for i := range passwd {
		passwd[i] = charset[rand.Intn(len(charset))]
	}

	return string(passwd)
}
