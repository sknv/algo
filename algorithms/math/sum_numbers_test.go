package math

import (
	"testing"
)

func TestSumGeneric(t *testing.T) {
	_, err := SumGeneric("1", 0, 3.)
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
	sum, err := SumGeneric(1, 0, 3.)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if sum != expected {
		t.Errorf("expect SumGeneric(...) to be [%f], got instead [%f]", expected, sum)
	}
}

func TestSumStringIntegers(t *testing.T) {
	_, err := SumStringIntegers("1d3", "25")
	if err == nil {
		t.Error("error expected, but none is returned")
	}

	_, err = SumStringIntegers("143", "25a")
	if err == nil {
		t.Error("error expected, but none is returned")
	}

	expected := "417"
	sum, err := SumStringIntegers("13", "404")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumStringIntegers(...) to be [%s], got instead [%s]", expected, sum)
	}

	expected = "6013623"
	sum, err = SumStringIntegers("219", "6013404")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumStringIntegers(...) to be [%s], got instead [%s]", expected, sum)
	}

	expected = "413"
	sum, err = SumStringIntegers("0", "413")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumStringIntegers(...) to be [%s], got instead [%s]", expected, sum)
	}

	expected = "0"
	sum, err = SumStringIntegers("0", "0")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumStringIntegers(...) to be [%s], got instead [%s]", expected, sum)
	}

	expected = "9223372036854775812"
	sum, err = SumStringIntegers("5", "9223372036854775807")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != expected {
		t.Errorf("expect SumStringIntegers(...) to be [%s], got instead [%s]", expected, sum)
	}
}
