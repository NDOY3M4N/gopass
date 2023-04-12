# 🔒 GoPass

![made-in-galsen](https://raw.githubusercontent.com/GalsenDev221/made.in.senegal/master/assets/badge.svg)

GoPass is a super simple CLI program for generating "random" passwords.

## Usage

After you've cloned the repo, you need to first define the options in the `main.go` file

```go
option := internal.Option{
  Length: 16,
  HasUppercase: true,
  HasLowercase: true,
  HasNumber: true,
  HasSymbol: true
}

pwd, score := internal.Generate(option)
```

And now, you can run the program with `go run .` or build it `go build` to generate passwords.
