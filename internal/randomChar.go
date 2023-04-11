package internal

import (
	"fmt"
	"math/rand"
)

var symbols = "!@#$%^&*()-_=+[]{}|;:,.<>/\\?~'\""

func RandomLowercase() string {
	n := rand.Intn(26)
	c := rune(n + 'a')

	return string(c)
}

func RandomUppercase() string {
	n := rand.Intn(26)
	c := rune(n + 'A')

	return string(c)
}

func RandomNumber() string {
	return fmt.Sprintf("%d", rand.Intn(10))
}

func RandomSymbol() string {
	c := symbols[rand.Intn(len(symbols))]

	return string(c)
}
