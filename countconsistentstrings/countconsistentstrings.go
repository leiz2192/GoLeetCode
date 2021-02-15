package countconsistentstrings

// countConsistentStrings https://leetcode-cn.com/problems/count-the-number-of-consistent-strings/
func countConsistentStrings(allowed string, words []string) (rst int) {
	b := [26]int{}
	for _, v := range allowed {
		b[v - 'a'] = 1
	}

	for _, word := range words {
		c := true
		for _, v := range word {
			if b[v - 'a'] != 1 {
				c = false
				break
			}
		}
		if c {
			rst++
		}
	}
	return rst
}
