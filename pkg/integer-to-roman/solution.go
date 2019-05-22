package integer_to_roman

import "strings"

// https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res, tmpValue := "", num
	for _, v := range values {
		if tmpValue != 0 && tmpValue < v {
			continue
		}

		if tmpValue == 0 {
			return res
		}

		nRepeatedRoman := tmpValue / v
		res += strings.Repeat(romanMap[v], nRepeatedRoman)
		tmpValue %= v
		if tmpValue == 0 {
			return res
		}
	}

	return res
}
