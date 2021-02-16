package subsets

import (
	"reflect"
	"testing"
)

func deepEqual(expected, actual [][]int) bool {
	if len(expected) != len(actual) {
		return false
	}

	size := len(expected)
	for i := 0; i < size; i++ {
		found := false
		oneExpected := expected[i]
		for j := 0; j < size; j++ {
			if lhs, rhs := len(oneExpected), len(actual[j]); lhs == rhs && lhs == 0 {
				found = true
				break
			}
			if reflect.DeepEqual(actual[j], oneExpected) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

func TestSubsets01(t *testing.T) {
	expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	nums := []int{1, 2, 3}
	if actual := subsets(nums); !deepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestSubsets02(t *testing.T) {
	nums := []int{0}
	expected := [][]int{{}, {0}}
	if actual := subsets(nums); !deepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
