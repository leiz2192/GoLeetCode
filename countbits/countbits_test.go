package countbits

import (
	"reflect"
	"testing"
)

func TestCountBits01(t *testing.T) {
	expected := []int{0, 1, 1}
	if actual := countBits(2); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestCountBits02(t *testing.T) {
	expected := []int{0, 1, 1, 2, 1, 2}
	if actual := countBits(5); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
