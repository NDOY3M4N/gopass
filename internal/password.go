package internal

import "math"

type Option struct {
	Length       int
	HasUppercase bool
	HasLowercase bool
	HasNumber    bool
	HasSymbol    bool
}

func NewOption() Option {
	return Option{Length: 8, HasLowercase: true}
}

func Generate(o Option) (pwd, entropy string) {
	var result string
	var possibleChars int

	if o.HasLowercase {
		possibleChars += 26
	}
	if o.HasUppercase {
		possibleChars += 26
	}
	if o.HasNumber {
		possibleChars += 10
	}
	if o.HasSymbol {
		possibleChars += len(symbols)
	}

	for i := 0; i < o.Length; i++ {
		// rand.Seed(time.Now().UnixNano())

		if len(result) < o.Length {
			if o.HasLowercase {
				result += RandomLowercase()
				continue
			}
			if o.HasUppercase {
				result += RandomUppercase()
				continue
			}
			if o.HasNumber {
				result += RandomNumber()
				continue
			}
			if o.HasSymbol {
				result += RandomSymbol()
				continue
			}
		}
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
