package arraypairsum

import "sort"

// ArrayPairSum https://leetcode-cn.com/problems/array-partition-i/
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)

	rst := 0
	for i, size := 0, len(nums); i < size; i += 2 {
		rst += nums[i]
	}
	return rst
}
