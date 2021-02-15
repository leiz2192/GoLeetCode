package generateparenthesis

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	expected := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"}
	if actual := GenerateParenthesis(3); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
