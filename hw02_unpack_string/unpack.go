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

	var previousSymbol rune
	for i, symbol := range str {
		// symbol is a rune

		if unicode.IsDigit(symbol) {
			if unicode.IsDigit(previousSymbol) || i == 0 {
				// "3abc" is invalid (starts with digit)
				// "aaa10b", "45" are invalid (digit before digit, eg. a number)
				return "", ErrInvalidString
			}

			if digitRuneToInt(symbol) == 0 && !unicode.IsDigit(previousSymbol) {
				// Do not write to builder if symbol is Zero (aaa0b -> aab)
				continue
			}
			// Repeat a previous symbol
			builder.WriteString(strings.Repeat(string(previousSymbol), digitRuneToInt(symbol)-1))
		} else {
			if utf8.RuneCountInString(str)-1 > i && rune(str[i+1]) == '0' {
				// Do not write to builder if next symbol is Zero (aaa0b -> aab)
				continue
			}
			// Write a symbol
			builder.WriteRune(symbol)
		}
		previousSymbol = symbol
	}

	return builder.String(), nil
}

func digitRuneToInt(r rune) int {
	return int(r - '0')
}
