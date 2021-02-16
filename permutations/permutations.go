package permutations

// https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	nsize := len(nums)
	rst := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == nsize {
			buf := make([]int, nsize)
			copy(buf, path)
			rst = append(rst, buf)
			return
		}

		for _, n := range nums {
			if v, ok := visited[n]; ok && v {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return rst
}
