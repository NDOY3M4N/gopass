package internal

type Options struct {
	length       int
	hasUppercase bool
	hasLowercase bool
	hasNumber    bool
	hasSymbol    bool
}

func Generate(o Options) string {
	var result string

	for i := 0; i < o.length; i++ {
		// rand.Seed(time.Now().UnixNano())

		if len(result) < o.length {
			if o.hasLowercase {
				result += RandomLowercase()
			}
			if o.hasUppercase {
				result += RandomUppercase()
			}
			if o.hasNumber {
				result += RandomNumber()
			}
			if o.hasSymbol {
				result += RandomSymbol()
			}
		}
	}

	return result
}
