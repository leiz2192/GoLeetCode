package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotate01(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if rotate(matrix); !reflect.DeepEqual(matrix, expected) {
		t.Errorf("expected: %v, actual: %v", expected, matrix)
	}
}

func TestRotate02(t *testing.T) {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	expected := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}
	if rotate(matrix); !reflect.DeepEqual(matrix, expected) {
		t.Errorf("expected: %v, actual: %v", expected, matrix)
	}
}

func TestRotate03(t *testing.T) {
	matrix := [][]int{{1}}
	expected := [][]int{{1}}
	if rotate(matrix); !reflect.DeepEqual(matrix, expected) {
		t.Errorf("expected: %v, actual: %v", expected, matrix)
	}
}

func TestRotate04(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}}
	expected := [][]int{{3, 1}, {4, 2}}
	if rotate(matrix); !reflect.DeepEqual(matrix, expected) {
		t.Errorf("expected: %v, actual: %v", expected, matrix)
	}
}
