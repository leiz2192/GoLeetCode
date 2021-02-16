package permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1}}
	if actual := permute(nums); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
