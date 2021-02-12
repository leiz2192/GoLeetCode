package increasingdecreasingstring

import "testing"

func TestSortString01(t *testing.T) {
	expected := "abccbaabccba"
	if actual := SortString("aaaabbbbcccc"); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSortString02(t *testing.T) {
	expected := "art"
	if actual := SortString("rat"); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSortString03(t *testing.T) {
	expected := "cdelotee"
	if actual := SortString("leetcode"); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
