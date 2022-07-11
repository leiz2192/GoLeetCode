package missingnumber

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := missingNumber(tc.nums)
			if got != tc.expected {
				t.Fatalf("expected: %d, got: %d", tc.expected, got)
			}
		})
	}
}
