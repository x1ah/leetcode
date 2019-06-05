package valid_parentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	if s == "" {
		return true
	}

	valueMap := map[byte]int{
		'[': 1,
		']': -1,
		'{': 2,
		'}': -2,
		'(': 3,
		')': -3,
	}

	var pairStack []int
	for i := range s {
		if len(pairStack) != 0 && pairStack[len(pairStack)-1]+valueMap[s[i]] == 0 {
			pairStack = pairStack[:len(pairStack)-1]
		} else {
			pairStack = append(pairStack, valueMap[s[i]])
		}
	}

	if len(pairStack) == 0 {
		return true
	}
	return false
}
