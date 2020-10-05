package math

import (
	"bytes"
	"fmt"
	"math"
)

func SumGeneric(numbers ...interface{}) (float64, error) {
	result := 0.
	for _, num := range numbers {
		switch val := num.(type) {
		case int:
			result += float64(val)
		case float64:
			result += val
		default:
			return 0, fmt.Errorf("invalid type [%T] provided, only int and float64 are allowed", val)
		}
	}
	return result, nil
}

func SumStringIntegers(a, b string) (string, error) {
	arr1, err := stringNumberToInt8Slice(a)
	if err != nil {
		return "", fmt.Errorf("failed to convert string number to int slice: %w", err)
	}

	arr2, err := stringNumberToInt8Slice(b)
	if err != nil {
		return "", fmt.Errorf("failed to convert string number to int slice: %w", err)
	}

	len1, len2 := len(arr1), len(arr2)
	maxLen := int(math.Max(float64(len1), float64(len2)))
	digits, carry := make([]int8, 0, maxLen), int8(0)
	for i, j := len1-1, len2-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x, y := int8(0), int8(0)
		if i >= 0 {
			x = arr1[i]
		}
		if j >= 0 {
			y = arr2[j]
		}

		sum := x + y + carry
		val := sum % 10
		carry = sum / 10
		digits = append(digits, val)
	}
	if carry > 0 {
		digits = append(digits, carry)
	}

	digits = reverseInt8Slice(digits)

	// Join the slice elements
	var buf bytes.Buffer
	for _, digit := range digits {
		num := digitInt8ToRune(digit)
		buf.WriteRune(num)
	}
	return buf.String(), nil
}

func stringNumberToInt8Slice(val string) ([]int8, error) {
	result := make([]int8, 0, len(val))
	for _, chr := range val {
		num := digitRuneToInt8(chr)
		if num < 0 || num > 9 {
			return nil, fmt.Errorf("invalid char [%v] provided in string [%v]", string(chr), val)
		}
		result = append(result, num)
	}
	return result, nil
}

func digitRuneToInt8(digit rune) int8 {
	return int8(digit - '0')
}

func digitInt8ToRune(digit int8) rune {
	return rune(digit + '0')
}

func reverseInt8Slice(arr []int8) []int8 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
