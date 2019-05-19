package zigzag_conversion

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	var res [][]byte

	currentRow, prefix := 0, 1
	for _, char := range s {
		if len(res) <= currentRow {
			tmp := []byte{byte(char)}
			res = append(res, tmp)
		} else {
			res[currentRow] = append(res[currentRow], byte(char))
		}
		if currentRow == numRows-1 {
			prefix = -1
		} else if currentRow == 0 {
			prefix = 1
		}
		currentRow += prefix
	}

	var bs []byte
	for _, lines := range res {
		for _, c := range lines {
			bs = append(bs, c)
		}
	}

	return string(bs[:])
}
