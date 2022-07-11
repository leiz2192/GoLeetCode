package missingnumber

// missingNumber https://leetcode.cn/problems/missing-number/
func missingNumber(nums []int) int {
	pos := make([]int, len(nums)+1)
	for _, n := range nums {
		pos[n] = 1
	}
	for i, p := range pos {
		if p == 0 {
			return i
		}
	}
	return 0
}