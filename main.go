package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NDOY3M4N/gopass/internal"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randLower := internal.RandomLowercase()
	randUpper := internal.RandomUppercase()
	randNumber := internal.RandomNumber()
	randSymbol := internal.RandomSymbol()

	fmt.Println("Random lowercase", randLower)
	fmt.Println("Random uppercase", randUpper)
	fmt.Println("Random number", randNumber)
	fmt.Println("Random symbol", randSymbol)
}
