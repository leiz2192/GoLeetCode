package leetcode55_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// canJump https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	var k int
	for i, n := range nums {
		if i > k {
			return false
		}
		if i+n > k {
			k = i + n
		}
	}
	return true
}

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums []int
		exp  bool
	}{
		{[]int{}, false},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0, 2, 3}, false},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := canJump(c.nums)
			assert.Equal(t, c.exp, got)
		})
	}
}
