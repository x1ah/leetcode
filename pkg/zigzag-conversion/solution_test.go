package zigzag_conversion

import "testing"

type zTestCase struct {
	input  string
	output string
	nrows  int
}

func TestZigzagConversion(t *testing.T) {
	testcases := []zTestCase{
		{input: "PAYPALISHIRING", output: "PINALSIGYAHRPI", nrows: 4},
		{input: "PAYPALISHIRING", output: "PAHNAPLSIIGYIR", nrows: 3},
	}

	for _, testcase := range testcases {
		output := convert(testcase.input, testcase.nrows)
		if output != testcase.output {
			t.Fatalf("%s expected %s, but got %s", testcase.input, testcase.output, output)
		}
	}
}
