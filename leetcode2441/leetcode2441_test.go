package leetcode2441_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// findMaxK https://leetcode.cn/problems/largest-positive-integer-that-exists-with-its-negative/
func findMaxK(nums []int) int {
	max := -1
	posInts := make(map[int]struct{})
	negInts := make(map[int]struct{})
	for _, n := range nums {
		if n > 0 {
			if _, ok := posInts[n]; !ok {
				posInts[n] = struct{}{}
			}
			if _, ok := negInts[-n]; ok && n > max {
				max = n
			}
			continue
		}

		if _, ok := negInts[n]; !ok {
			negInts[n] = struct{}{}
		}

		if _, ok := posInts[-n]; ok && -n > max {
			max = -n
		}
	}
	return max
}

func TestFindMaxK(t *testing.T) {
	cases := []struct {
		nums []int
		exp  int
	}{
		{nums: []int{}, exp: -1},
		{nums: []int{-10, 8, 6, 7, -2, -3}, exp: -1},
		{nums: []int{-1, 2, -3, 3}, exp: 3},
		{nums: []int{-1, 10, 6, 7, -7, 1}, exp: 7},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := findMaxK(c.nums)
			assert.Equal(t, c.exp, got)
		})
	}
}
