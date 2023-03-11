package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	result := ""
	for index := range input {
		if input[index] >= 48 && input[index] <= 57 {
			if index-1 < 0 || input[index-1] >= 48 && input[index-1] <= 57 {
				return "", ErrInvalidString
			}
			count := int(input[index] - '0')
			if count > 0 {
				count = count - 1
				result = result + strings.Repeat(string(input[index-1]), count)
			} else if count == 0 {
				result = string(input[:index-1])
			}
		} else {
			result = result + string(input[index])
		}
	}
	return result, nil
}
