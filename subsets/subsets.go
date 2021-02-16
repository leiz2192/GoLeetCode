package subsets

// https://leetcode-cn.com/problems/subsets/
func subsets(nums []int) (ans [][]int) {
	set := []int{}
	nsize := len(nums)
	var dfs func(index int)
	dfs = func(index int) {
		if index == nsize {
			ans = append(ans, append([]int(nil), set...))
			return
		}

		set = append(set, nums[index])
		dfs(index + 1)
		set = set[:len(set)-1]
		dfs(index + 1)
	}
	dfs(0)
	return
}
