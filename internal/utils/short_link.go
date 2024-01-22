package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateShortCode() string {
	rand.NewSource(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return time.Now().Format("20060102") + hex.EncodeToString([]byte(string(b)))
}
