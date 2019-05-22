package integer_to_roman

import "testing"

func TestIntToRoman(t *testing.T) {
	testcases := map[int]string{
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}

	for num, roman := range testcases {
		res := intToRoman(num)
		if res != roman {
			t.Fatalf("%d expected %s, bug got %s", num, roman, res)
		}
	}
}
