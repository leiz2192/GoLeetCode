package buildanarray_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// buildArray https://leetcode.cn/problems/build-an-array-with-stack-operations/
func buildArray(target []int, n int) []string {
	rst := make([]string, 0, n)
	pre := 0
	for _, num := range target {
		for i := 1; i < num-pre; i++ {
			rst = append(rst, "Push", "Pop")
		}
		rst = append(rst, "Push")
		pre = num
	}
	return rst
}

func TestBuildArray(t *testing.T) {
	testCases := []struct {
		target []int
		n      int
		exp    []string
	}{
		{[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
		{[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
		{[]int{1, 2}, 4, []string{"Push", "Push"}},
	}

	for index, tc := range testCases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			got := buildArray(tc.target, tc.n)
			assert.Equal(t, tc.exp, got)
		})
	}
}
