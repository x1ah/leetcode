package roman_to_integer

import "testing"

func TestRomanToInteger(t *testing.T) {
	testcases := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for roman, value := range testcases {
		if romanToInt(roman) != value {
			t.Fatalf("%s expected %d, but got %d\n", roman, value, romanToInt(roman))
		}
	}
}
