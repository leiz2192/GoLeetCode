package countconsistentstrings

import "testing"

func TestCountConsistentStrings01(t *testing.T) {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	expected := 2
	if actual := countConsistentStrings(allowed, words); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCountConsistentStrings02(t *testing.T) {
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	expected := 7
	if actual := countConsistentStrings(allowed, words); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCountConsistentStrings03(t *testing.T) {
	allowed := "cad"
	words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	expected := 4
	if actual := countConsistentStrings(allowed, words); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
