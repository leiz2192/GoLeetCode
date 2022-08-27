package reversewords_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// reverseWords https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	reverse := func(b []byte) {
		lhs, rhs := 0, len(b)-1
		for lhs < rhs {
			b[lhs], b[rhs] = b[rhs], b[lhs]
			lhs++
			rhs--
		}
	}
	b := []byte(s)
	reverse(b)

	// words := make([]byte, 0, len(b))
	words := b[0:0:len(b)]
	lhs, rhs := 0, 0
	for lhs < len(b) {
		if b[lhs] == ' ' {
			lhs++
			continue
		}
		rhs = lhs + 1
		for rhs < len(b) {
			if b[rhs] != ' ' {
				rhs++
				continue
			}
			reverse(b[lhs:rhs])
			words = append(words, b[lhs:rhs]...)
			words = append(words, ' ')
			break
		}
		if rhs == len(b) {
			reverse(b[lhs:rhs])
			words = append(words, b[lhs:rhs]...)
			words = append(words, ' ')
			break
		}
		lhs = rhs + 1
	}
	return string(words[0:len(words)-1])
}

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		s   string
		exp string
	}{
		{s: "the sky is blue", exp: "blue is sky the"},
		{s: "  hello world  ", exp: "world hello"},
		{s: "a good   example", exp: "example good a"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := reverseWords(tc.s)
			assert.Equal(t, tc.exp, got)
		})
	}
}
