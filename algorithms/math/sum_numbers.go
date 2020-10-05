package math

import (
	"fmt"
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
			return 0, fmt.Errorf("invalid type %T provided, only int and float64 are allowed", val)
		}
	}
	return result, nil
}
