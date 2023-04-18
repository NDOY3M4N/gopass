package password

import (
	"math"
	"math/rand"
)

// NOTE: should I create a `Password` struct?
// That way, I would have a generate & score method
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
		// NOTE: this isn't what I want btw but I'll keep it for now
		// I wanted to insert the characters in a random order
		randomIndex := rand.Intn(len(fns))
		result += fns[randomIndex]()
	}

	score := calculateScore(o.Length, possibleChars)

	return result, score
}

const (
	score1 = "Too Weak!"
	score2 = "Weak"
	score3 = "Medium"
	score4 = "Strong"
)

func calculateScore(length, possibleChars int) string {
	entropy := math.Log2(math.Pow(float64(possibleChars), float64(length)))

	if entropy < 38 {
		return score1
	} else if entropy < 46 {
		return score2
	} else if entropy < 55 {
		return score3
	} else {
		return score4
	}
}
