package utils

import "math/rand"

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(r int) string {
	bits := []rune{}
	k := len(alphabet)

	for i := 0; i < r; i++ {
		index := rand.Intn(k)
		bits = append(bits, rune(alphabet[index]))
	}
	return string(bits)
}

func RandomEmail() string {
	return RandomString(10) + "@example.com"
}