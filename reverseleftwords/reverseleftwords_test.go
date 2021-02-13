package reverseleftwords

import "testing"

func TestReverseLeftWords01(t *testing.T) {
	expected := "cdefgab"
	if actual := ReverseLeftWords("abcdefg", 2); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestReverseLeftWords02(t *testing.T) {
	expected := "umghlrlose"
	if actual := ReverseLeftWords("lrloseumgh", 6); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
