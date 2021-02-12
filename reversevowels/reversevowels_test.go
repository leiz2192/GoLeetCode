package reversevowels

import "testing"

func TestReverseVowels01(t *testing.T) {
	expected := "holle"
	if actual := ReverseVowels("hello"); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestReverseVowels02(t *testing.T) {
	expected := "leotcede"
	if actual := ReverseVowels("leetcode"); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
