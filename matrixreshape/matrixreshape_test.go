package matrixreshape

import (
	"reflect"
	"testing"
)

func TestMatrixReshape01(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	expected := [][]int{{1, 2, 3, 4}}
	if actual := matrixReshape(nums, 1, 4); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestMatrixReshape02(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	expected := [][]int{{1, 2}, {3, 4}}
	if actual := matrixReshape(nums, 2, 4); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
