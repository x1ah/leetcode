package letter_combinations_of_a_phone_number

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	backtrack("", digits)
	return res
}

func backtrack(prefix, nextDigits string) {
	if len(nextDigits) == 0 {
		res = append(res, prefix)
		return
	}
	digitMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	chars := digitMap[nextDigits[0]]
	for _, char := range chars {
		backtrack(prefix+char, nextDigits[1:])
	}
}
