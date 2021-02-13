package finddisappearednumbers

// FindDisappearedNumbers https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	rst := []int{}
	for i, v := range nums {
		if v <= n {
			rst = append(rst, i+1)
		}
	}
	return rst
}
