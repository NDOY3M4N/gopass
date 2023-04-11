package internal

import (
	"regexp"
	"testing"
)

func TestPassword(t *testing.T) {
	// Same character type
	t.Run("generate an 8-long password with lowercases only", func(t *testing.T) {
		options := Options{length: 8, hasLowercase: true}
		got := Generate(options)

		assertCorrectPassword(t, got, `^[a-z]{8}$`, "lowercase only")
	})

	t.Run("generate an 9-long password with uppercases only", func(t *testing.T) {
		options := Options{length: 9, hasUppercase: true}
		got := Generate(options)

		assertCorrectPassword(t, got, `^[A-Z]{9}$`, "uppercase only")
	})

	t.Run("generate an 10-long password with numbers only", func(t *testing.T) {
		options := Options{length: 10, hasNumber: true}
		got := Generate(options)

		assertCorrectPassword(t, got, `^[0-9]{10}$`, "number only")
	})

	t.Run("generate an 11-long password with symbols only", func(t *testing.T) {
		options := Options{length: 11, hasSymbol: true}
		got := Generate(options)

		assertCorrectPassword(t, got, `^[_\W]{11}$`, "symbol only")
	})

	// Mixed characters
	t.Run("generate an 8-long password with uppercases and lowercase", func(t *testing.T) {
		options := Options{length: 8, hasLowercase: true, hasUppercase: true}
		got := Generate(options)

		// TODO: need to fix this regex
		assertCorrectPassword(t, got, "^[a-zA-Z]{8}$", "uppercase and lowercase")
	})

	t.Run("generate an 16-long password with uppercases and numbers", func(t *testing.T) {
		options := Options{length: 16, hasNumber: true, hasUppercase: true}
		got := Generate(options)

		// TODO: need to fix this regex
		assertCorrectPassword(t, got, `^[A-Z0-9]{16}$`, "uppercase and lowercase")
		// (?=.*\d)((?=.*\W)|(?=.*_))^[^ ]{8}$
	})
}

func assertCorrectPassword(t testing.TB, s, pattern, charType string) {
	t.Helper()

	match, err := regexp.MatchString(pattern, s)

	if err != nil {
		panic(err)
	}

	if !match {
		t.Errorf("%q need to have %s", s, charType)
	}
}
