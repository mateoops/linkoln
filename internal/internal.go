package internal

import (
	"math/rand"
)

func GenerateShortID(length int) string {
	const defaultChars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = defaultChars[rand.Intn(len(defaultChars))]
	}
	return string(result)
}
