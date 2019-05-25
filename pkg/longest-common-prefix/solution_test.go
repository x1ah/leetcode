package longest_common_prefix

import (
	"testing"
)

func TestLCP(t *testing.T) {
	testcases := map[string][2]string{
		"abcd": [2]string{"abcde", "abcddddd"},
		"ab":   [2]string{"abecde", "abcddddd"},
		"":     [2]string{"acbdd", "lkj;"},
	}

	for answer, s := range testcases {
		res := lcp(s[0], s[1])
		if answer != res {
			t.Fatalf("%v excepted %s, but got %s", s, answer, res)
		}
	}
}

func TestSolution(t *testing.T) {
	testcases := map[string][]string{
		"fl": []string{"flower", "flow", "flight"},
		//	"": []string{"dog","racecar","car"},
		"": []string{"ca", "a"},
	}

	for answer, s := range testcases {
		res := longestCommonPrefix2(s)
		res2 := longestCommonPrefix(s)
		if res != answer || res2 != answer {
			t.Fatalf("%v excepted %s, bug got %s", s, answer, res)
		}
	}
}
