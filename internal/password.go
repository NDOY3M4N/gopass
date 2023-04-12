package internal

import (
	"math"
	"math/rand"
)

type Option struct {
	Length       int
	HasUppercase bool
	HasLowercase bool
	HasNumber    bool
	HasSymbol    bool
}

// NOTE: Do I need this bad boi?
func NewOption() Option {
	return Option{Length: 8, HasLowercase: true}
}

func Generate(o Option) (pwd, entropy string) {
	var result string
	var possibleChars int
	var fns []func() string

	if o.HasLowercase {
		possibleChars += 26
		fns = append(fns, RandomLowercase)
	}
	if o.HasUppercase {
		possibleChars += 26
		fns = append(fns, RandomUppercase)
	}
	if o.HasNumber {
		possibleChars += 10
		fns = append(fns, RandomNumber)
	}
	if o.HasSymbol {
		possibleChars += len(symbols)
		fns = append(fns, RandomSymbol)
	}

	for i := 0; i < o.Length; i++ {
		// TODO: this isn't what I want.
		randomIndex := rand.Intn(len(fns))
		result += fns[randomIndex]()
	}

	score := calculateScore(o.Length, possibleChars)

	return result, score
}

func calculateScore(length, possibleChars int) string {
	entropy := math.Log2(math.Pow(float64(possibleChars), float64(length)))

	// TODO: change this scoring system
	if entropy < 28 {
		return "Too Weak!"
	} else if entropy < 36 {
		return "Weak"
	} else if entropy < 60 {
		return "Medium"
	} else {
		return "Strong"
	}
}
