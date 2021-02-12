package increasingdecreasingstring

// SortString https://leetcode-cn.com/problems/increasing-decreasing-string/
func SortString(s string) string {
	b := [26]int{}
	for _, c := range s {
		b[c-'a']++
	}

	rst := make([]byte, len(s))
	for k, size := 0, len(b); k < len(s); {
		for i := 0; i < size; i++ {
			if b[i] > 0 {
				rst[k] = byte(i) + 'a'
				k++
				b[i]--
			}
		}

		for j := size - 1; j >= 0; j-- {
			if b[j] > 0 {
				rst[k] = byte(j) + 'a'
				k++
				b[j]--
			}
		}
	}
	return string(rst)
}
