package finddisappearednumbers

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers01(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}
	if actual := FindDisappearedNumbers(nums); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
