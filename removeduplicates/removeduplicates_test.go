package removeduplicates_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// removeDuplicates from https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lhs, rhs := 0, 0
	for rhs < len(nums) {
		if nums[lhs] != nums[rhs] {
			lhs++
			nums[lhs] = nums[rhs]
		}
		rhs++
	}

	return lhs + 1
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums    []int
		expLen  int
		expNums []int
	}{
		{nums: []int{1, 1, 2}, expLen: 2, expNums: []int{1, 2}},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expLen: 5, expNums: []int{0, 1, 2, 3, 4}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := removeDuplicates(tc.nums)
			assert.Equal(t, tc.expLen, got)
			for i := 0; i < tc.expLen; i++ {
				assert.Equal(t, tc.expNums[i], tc.nums[i])
			}
		})
	}
}
