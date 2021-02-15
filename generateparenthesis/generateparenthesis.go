package generateparenthesis

import "fmt"

// GenerateParenthesis https://leetcode-cn.com/problems/bracket-lcci/
func GenerateParenthesis(n int) []string {
	rst := []string{}
	pathsize := 2 * n

	var dfs func(path string, lhs int, rhs int)
	dfs = func(path string, lhs int, rhs int) {
		if pathsize == len(path) {
			rst = append(rst, path)
			return
		}

		if lhs > 0 {
			dfs(fmt.Sprintf("%s%s", path, "("), lhs-1, rhs)
		}

		if rhs > lhs {
			dfs(fmt.Sprintf("%s%s", path, ")"), lhs, rhs-1)
		}
	}

	dfs("", n, n)
	return rst
}
