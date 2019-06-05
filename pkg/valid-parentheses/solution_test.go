package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testcases := map[string]bool{
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
	}

	for s, answer := range testcases {
		if isValid(s) != answer {
			t.Fatalf("%s failed", s)
		}
	}
}
