package leetcode1016_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// queryString https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n/
func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		v := strconv.FormatInt(int64(i), 2)
		if !strings.Contains(s, v) {
			return false
		}
	}
	return true
}

func TestQueryString(t *testing.T) {
	cases := []struct {
		s   string
		n   int
		exp bool
	}{
		{"110101011011000011011111000000", 15, false},
		{"", 3, false},
		{"0110", 3, true},
		{"0110", 4, false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := queryString(c.s, c.n)
			assert.Equal(t, c.exp, got)
		})
	}
}
