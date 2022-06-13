package hw02unpackstring

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (output string, err error) {
	for i, r := range input {
		if unicode.IsDigit(r) {
			if i == 0 || unicode.IsDigit(rune(input[i-1])) {
				return "", ErrInvalidString
			}
			n := int(r - '0')
			for j := 1; j < n; j++ {
				output += string(input[i-1])
			}
			if n == 0 {
				length := utf8.RuneCountInString(output)
				output = output[:length-1]
			}
		}
		if unicode.IsLetter(r) {
			output += string(input[i])
		}
	}
	return output, nil
}
