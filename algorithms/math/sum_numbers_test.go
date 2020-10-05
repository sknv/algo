package math

import (
	"testing"
)

func TestSumGeneric(t *testing.T) {
	var (
		sum float64
		err error
	)

	_, err = SumGeneric("1", 0, 3.)
	if err == nil {
		t.Error("error expected, but none is returned")
	}

	_, err = SumGeneric(int32(1), 0, 3.)
	if err == nil {
		t.Error("error expected, but none is returned")
	}

	_, err = SumGeneric(1, int64(0), 3.)
	if err == nil {
		t.Error("error expected, but none is returned")
	}

	expected := 4.
	sum, err = SumGeneric(1, 0, 3.)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumGeneric(...) to be: %f, got instead: %f", expected, sum)
	}
}
