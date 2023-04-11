package internal

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

func Generate(o Option) string {
	var result string

	for i := 0; i < o.Length; i++ {
		// rand.Seed(time.Now().UnixNano())

		if len(result) < o.Length {
			if o.HasLowercase {
				result += RandomLowercase()
			}
			if o.HasUppercase {
				result += RandomUppercase()
			}
			if o.HasNumber {
				result += RandomNumber()
			}
			if o.HasSymbol {
				result += RandomSymbol()
			}
		}
	}

	return result
}
