package password

import (
	"regexp"
	"testing"
)

func TestRandomCharacter(t *testing.T) {
	t.Run("get a random uppercase letter", func(t *testing.T) {
		s := RandomUppercase()
		assertCorrectRegex(t, s, "[A-Z]", "uppercase")
	})

	t.Run("get a random lowercase letter", func(t *testing.T) {
		s := RandomLowercase()
		assertCorrectRegex(t, s, "[a-z]", "lowercase")
	})

	t.Run("get a random number ", func(t *testing.T) {
		s := RandomNumber()
		assertCorrectRegex(t, s, "[0-9]", "number")
	})

	t.Run("get a random symbol", func(t *testing.T) {
		s := RandomSymbol()
		assertCorrectRegex(t, s, "[_\\W]", "symbol only password")
	})
}

func assertCorrectRegex(t testing.TB, character, pattern, charType string) {
	t.Helper()

	matched, err := regexp.MatchString(pattern, character)

	if err != nil {
		panic(err)
	}

	if !matched {
		t.Errorf("the random character is not a %s", charType)
	}
}
