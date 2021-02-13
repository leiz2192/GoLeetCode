package sumnums

import "testing"

func TestSumNums01(t *testing.T) {
	expected := 6
	if actual := SumNums(3); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestSumNums02(t *testing.T) {
	expected := 45
	if actual := SumNums(9); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
