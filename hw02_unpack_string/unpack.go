package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder

	builder.WriteString("")
	_, lastSymbolSize := utf8.DecodeLastRuneInString(str)

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
			builder.WriteString(strings.Repeat(string(previousSymbol), digitRuneToInt(symbol)))

		case len(str)-lastSymbolSize == pos:
			if !unicode.IsDigit(previousSymbol) {
				builder.WriteRune(previousSymbol)
			}
			builder.WriteRune(symbol)

		case !unicode.IsDigit(previousSymbol) && pos != 0:
			builder.WriteRune(previousSymbol)
		}
		previousSymbol = symbol
	}

	return builder.String(), nil
}

func digitRuneToInt(r rune) int {
	return int(r - '0')
}
