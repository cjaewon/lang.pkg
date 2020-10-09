package lib

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

// GenerateRandKey : Generate Random Key
func GenerateRandKey(n int) string {
	k := make([]byte, n)
	for i := range k {
		k[i] = letters[rand.Intn(len(letters))]
	}

	return string(k)
}
