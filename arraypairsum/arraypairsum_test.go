package arraypairsum

import "testing"

func TestArrayPairSum01(t *testing.T) {
	nums := []int{1, 4, 3, 2}
	expected := 4
	if actual := ArrayPairSum(nums); actual != expected {
		t.Errorf("expected %d, actual: %d", expected, actual)
	}
}

func TestArrayPairSum02(t *testing.T) {
	nums := []int{6, 2, 6, 5, 1, 2}
	expected := 9
	if actual := ArrayPairSum(nums); actual != expected {
		t.Errorf("expected %d, actual: %d", expected, actual)
	}
}
