package main

import (
	"fmt"
	"github.com/NDOY3M4N/gopass/internal"
)

func main() {
	option := internal.Option{
		Length:       8,
		HasUppercase: true,
		// HasLowercase: true,
		// HasNumber: true,
		// HasSymbol: true,
	}

	pwd, score := internal.Generate(option)

	fmt.Println("Welcome to my humble cli app")
	fmt.Printf("The generated password is %q and the score is %v", pwd, score)
}
