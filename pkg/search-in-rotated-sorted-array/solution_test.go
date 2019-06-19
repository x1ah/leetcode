package search_in_rotated_sorted_array

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	answers := map[int][]int{
		-1: {9, 4, 5, 6, 1, 2, 3},
		4:  {0, 4, 5, 6, 7, 0, 1, 2},
	}

	for answer, v := range answers {
		res := search(v[1:], v[0])
		if answer != res {
			t.Fatalf("%v excepted %v, got %v", v, answer, res)
		}
	}
}
