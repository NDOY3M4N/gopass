package main

import (
	"fmt"
	"github.com/NDOY3M4N/gopass/internal"
)

func main() {
	option := internal.Option{
		Length:       16,
		HasUppercase: true,
		HasLowercase: true,
		HasNumber:    true,
		HasSymbol:    true,
	}

	pwd := internal.Generate(option)

	fmt.Println("Welcome to my humble cli app")
	fmt.Println("The generated password is", pwd)
}
