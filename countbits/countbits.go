package countbits

// https://leetcode-cn.com/problems/counting-bits/
func countBits(num int) []int {
	rst := make([]int, num+1)
	rst[0] = 0
	for i := 1; i <= num; i++ {
		rst[i] = rst[i/2] + (i & 1)
	}
	return rst
}
