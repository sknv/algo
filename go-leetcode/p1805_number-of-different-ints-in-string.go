package leetcode

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/number-of-different-integers-in-a-string/

func numDifferentIntegers(word string) int {
	replaced := replaceNonDigitsWithWhitespace(word)
	splitted := strings.Split(replaced, " ")

	numbers := make(map[string]struct{}, len(splitted))
	for _, number := range splitted {
		if number == "" {
			continue
		}

		trimmed := strings.TrimLeft(number, "0")
		numbers[trimmed] = struct{}{}
	}

	return len(numbers)
}

func replaceNonDigitsWithWhitespace(input string) string {
	runes := []rune(input)
	for i, r := range runes {
		if !unicode.IsDigit(r) {
			runes[i] = ' '
		}
	}

	return string(runes)
}
