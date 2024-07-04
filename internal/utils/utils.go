package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

// GenerateShortKey generates a random short key with the given length using the provided character set.
func GenerateShortKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
