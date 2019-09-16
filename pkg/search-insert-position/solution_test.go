package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {

	testcases := [][]int{
		{1, 2, 1, 3, 5, 6},
		{1, 2, 1, 3},
		{2, 5, 1, 3, 5, 6},
		{4, 7, 1, 3, 5, 6},
		{0, 0, 1, 3, 5, 6},
	}

	for _, testcase := range testcases {
		res := searchInsert(testcase[2:], testcase[1])
		if res != testcase[0] {
			t.Fatalf("%v, %v expected %v, but got %v", testcase[2:], testcase[1], testcase[0], res)
		}
	}
}
