package rotateimage

// https://leetcode-cn.com/problems/rotate-image/
func rotate(matrix [][]int) {
	n := len(matrix)
	buf := make([][]int, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, v := range row {
			buf[j][n-i-1] = v
		}
	}
	copy(matrix, buf)
}
