package sumnums

// SumNums https://leetcode-cn.com/problems/qiu-12n-lcof/
func SumNums(n int) int {
	rst := 0
	var sum func(int) bool
	sum = func(num int) bool {
		rst += num
		return num > 0 && sum(num-1)
	}
	_ = sum(n)
	return rst
}
