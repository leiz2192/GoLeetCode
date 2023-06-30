package leetcode2490_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isCircularSentence https://leetcode.cn/problems/circular-sentence/
func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	if len(words) == 1 {
		return sentence[0] == sentence[len(sentence)-1]
	}
	for i := 1; i < len(words); i++ {
		preWordLen := len(words[i-1])
		if words[i-1][preWordLen-1] != words[i][0] {
			return false
		}
	}
	lastWordLen := len(words[len(words)-1])
	return words[0][0] == words[len(words)-1][lastWordLen-1]
}

func TestIsCircularSentence(t *testing.T) {
	testCases := []struct {
		sentence string
		exp      bool
	}{
		{"leetcode exercises sound delightful", true},
		{"eetcode", true},
		{"Leetcode is cool", false},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := isCircularSentence(tc.sentence)
			assert.Equal(t, tc.exp, got)
		})
	}
}
