package roman_to_integer

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	valueMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length, res := len(s), 0
	cursor := 0
	for cursor < length {
		if (cursor != length-1) && (valueMap[s[cursor]] < valueMap[s[cursor+1]]) {
			res = res + valueMap[s[cursor+1]] - valueMap[s[cursor]]
			cursor += 2
		} else {
			res += valueMap[s[cursor]]
			cursor += 1
		}
	}

	return res
}
