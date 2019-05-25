package longest_common_prefix

// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	}

	newLcp := lcp(strs[0], strs[1])
	strs = append(strs[2:], newLcp)
	return longestCommonPrefix(strs)
}

// 最长公共前缀（二维数组、动态规划）
func lcp(a, b string) string {
	aLen, bLen := len(a), len(b)
	if aLen == 0 || bLen == 0 {
		return ""
	}

	if a[0] != b[0] {
		return ""
	}

	// 最长子串长度、最长子串在 a 字符串的 end cursor
	maxLength, aCursor := 0, 0
	var res [][]int
	for i := 0; i < aLen; i++ {
		res = append(res, make([]int, bLen))
	}

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if a[i] == b[j] && i == j {
				if i == 0 || j == 0 {
					res[i][j] = 1
				} else {
					res[i][j] = res[i-1][j-1] + 1
				}

				if res[i][j] >= maxLength {
					maxLength = res[i][j]
					aCursor = i
				}
			} else {
				res[i][j] = 0
			}
		}
	}

	if maxLength == 0 {
		return ""
	}
	return string(a[aCursor-maxLength+1 : aCursor+1])
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i >= len(str) || str[i] != strs[0][i] {
				return string(str[:i])
			}
		}
	}
	return string(strs[0])
}
