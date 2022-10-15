package sortanarray_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// sortArray https://leetcode.cn/problems/sort-an-array/
func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	lhs, rhs := start, end
	randomIndex := rand.Intn(rhs-lhs+1) + lhs
	x := nums[randomIndex]
	nums[randomIndex], nums[lhs] = nums[lhs], nums[randomIndex]

	for lhs < rhs {
		for lhs < rhs && nums[rhs] >= x {
			rhs--
		}
		if lhs < rhs {
			nums[lhs] = nums[rhs]
			lhs++
		}
		for lhs < rhs && nums[lhs] < x {
			lhs++
		}
		if lhs < rhs {
			nums[rhs] = nums[lhs]
			rhs--
		}
	}
	nums[lhs] = x
	quickSort(nums, start, lhs-1)
	quickSort(nums, lhs+1, end)
}

func TestSortArray(t *testing.T) {
	testCases := []struct {
		nums []int
		exp  []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := sortArray(tc.nums)
			assert.Equal(t, tc.exp, got)
		})
	}
}
