package minimumwindowsubstring_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// minWindow https://leetcode.cn/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	var left, right, validCnt, start int
	end := math.MaxInt
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				validCnt++
			}
		}

		for validCnt == len(need) {
			if right-left < end-start {
				start, end = left, right
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					validCnt--
				}
				window[d]--
			}
		}
	}
	if end < math.MaxInt {
		return s[start:end]
	}
	return ""
}

func TestMinWidow(t *testing.T) {
	for i, tc := range []struct {
		s   string
		t   string
		exp string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := minWindow(tc.s, tc.t)
			assert.Equal(t, tc.exp, got)
		})
	}
}
