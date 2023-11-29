package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomID() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(b)
}
