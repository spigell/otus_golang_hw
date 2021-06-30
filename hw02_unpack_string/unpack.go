package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	var builder strings.Builder

	var previousSymbol rune
	for pos, symbol := range str {
		// symbol is a rune

		switch {
		case unicode.IsDigit(symbol):
			if unicode.IsDigit(previousSymbol) || pos == 0 {
				// "3abc" is invalid (starts with digit)
				// "aaa10b", "45" are invalid (digit before digit, eg. a number)
				return "", ErrInvalidString
			}

			// Repeat a previous symbol
			builder.WriteString(strings.Repeat(string(previousSymbol), int(symbol-'0')))

		case !unicode.IsDigit(previousSymbol) && pos != 0:
			builder.WriteRune(previousSymbol)
		}
		previousSymbol = symbol
	}

	if !unicode.IsDigit(previousSymbol) {
		builder.WriteString(string(previousSymbol))
	}

	return builder.String(), nil
}
