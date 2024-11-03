package main

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(n int) string {
	password := make([]byte, n)
	for i := range password {
		password[i] = byte(rand.Int() % len(charset))
	}
	return string(password)
}

func main() {
	println(generatePassword(10)) // Пример генерации пароля длиной 10 символов
}
