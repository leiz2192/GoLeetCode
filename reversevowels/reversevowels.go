package reversevowels

var vowels = map[byte]int{'a': 1, 'o': 1, 'e': 1, 'i': 1, 'u': 1, 'A': 1, 'O': 1, 'E': 1, 'I': 1, 'U': 1}

// ReverseVowels https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
func ReverseVowels(s string) string {
	b := []byte(s)
	lhs, rhs := 0, len(b)-1
	for lhs < rhs {
		for !isVowels(b[lhs]) && lhs < rhs {
			lhs++
		}
		for !isVowels(b[rhs]) && lhs < rhs {
			rhs--
		}

		b[lhs], b[rhs] = b[rhs], b[lhs]
		lhs++
		rhs--
	}
	return string(b)
}

func isVowels(b byte) bool {
	_, ok := vowels[b]
	return ok
}
