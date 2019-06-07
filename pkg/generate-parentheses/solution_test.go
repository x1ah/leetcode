package generate_parentheses

import "testing"

func TestGenerateParentheses(t *testing.T) {
	res := generateParenthesis(3)
	if len(res) != 5 {
		t.Fatalf("fail")
	}

	answers := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	for i, answer := range answers {
		if res[i] != answer {
			t.Fatalf("%v != %v, failed", res[i], answer)
		}
	}
}
