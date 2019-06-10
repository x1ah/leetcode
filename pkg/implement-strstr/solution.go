package implement_strstr

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == len(haystack) {
		if needle == haystack {
			return 0
		} else {
			return -1
		}
	}

	needleLength := len(needle)
	for i := 0; i < len(haystack)-needleLength+1; i++ {
		if string(haystack[i:i+needleLength]) == needle {
			return i
		}
	}

	return -1
}
