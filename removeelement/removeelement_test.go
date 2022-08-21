package removeelement_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// removeElement https://leetcode.cn/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	lhs, rhs := 0, 0
	for rhs < len(nums) {
		if nums[rhs] != val {
			nums[lhs] = nums[rhs]
			lhs++
		}
		rhs++
	}

	return lhs
}

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums    []int
		val     int
		expLen  int
		expNums []int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expLen: 2, expNums: []int{2, 2}},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expLen: 5, expNums: []int{0, 1, 3, 0, 4}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := removeElement(tc.nums, tc.val)
			assert.Equal(t, tc.expLen, got)
			for i := 0; i < got; i++ {
				assert.Equal(t, tc.expNums[i], tc.nums[i], fmt.Sprintf("got: %v, exp: %v", tc.nums, tc.expNums))
			}
		})
	}
}
