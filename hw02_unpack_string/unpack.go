package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var result strings.Builder
	var previousRune rune
	var numCount int
	var isPrevDigit bool

	for _, runeChunk := range inputString {
		// checking for two digits located side by side, or not digit nor letter
		if (unicode.IsDigit(runeChunk) && isPrevDigit) || (!unicode.IsDigit(runeChunk) && !unicode.IsLetter(runeChunk)) {
			return "", ErrInvalidString
		}
		if unicode.IsLetter(runeChunk) {
			result.WriteRune(runeChunk)
			previousRune = runeChunk
			isPrevDigit = false
		}
		if unicode.IsDigit(runeChunk) {
			numCount, _ = strconv.Atoi(string(runeChunk))
			if numCount == 0 {
				// delete previous rune if 0 occure after that
				resultStr := result.String()
				result.Reset()
				result.WriteString(resultStr[:len(resultStr)-1])
			} else {
				if previousRune == 0 {
					return "", ErrInvalidString
				}
				result.WriteString(strings.Repeat(string(previousRune), numCount-1))
			}
			isPrevDigit = true
		}
	}
	return result.String(), nil
}
