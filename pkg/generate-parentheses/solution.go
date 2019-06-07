package generate_parentheses

var res []string

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	res = []string{}
	backtrack("", 0, 0, n)
	return res
}

func backtrack(s string, left int, right int, n int) {
	if len(s) == n*2 {
		res = append(res, s)
		return
	}

	if left < n {
		backtrack(s+"(", left+1, right, n)
	}

	// 只有右括号不足时才添加，保证添加后的字符串时有效的
	if right < left {
		backtrack(s+")", left, right+1, n)
	}
}
