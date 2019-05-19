package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {

	testcases := map[int][]int{
		49: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		36: []int{2, 3, 10, 5, 7, 8, 9},
	}

	for answer, heights := range testcases {
		res := maxArea(heights)
		if res != answer {
			t.Fatalf("%v expected %d, but got %d\n", heights, answer, res)
		}
	}
}
