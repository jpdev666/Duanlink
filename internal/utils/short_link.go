package utils

import (
	"math/rand"
	"strings"
	"time"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShortCode() string {
	rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.Int63n(62 * 62 * 62 * 62 * 62 * 62)
	return encodeBase62(randomNumber)
}

func encodeBase62(number int64) string {
	if number == 0 {
		return string(base62Chars[0])
	}
	var encodedBuilder strings.Builder
	base := int64(len(base62Chars))
	for number > 0 {
		remainder := number % base
		number /= base
		encodedBuilder.WriteByte(base62Chars[remainder])
	}
	return encodedBuilder.String()
}
