package three_sum_closest

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	testcases := map[int][]int{
		0: {0, 0, 0, 0},
		2: {2, -1, 1, 2, 4},
	}

	for answer, nums := range testcases {
		res := threeSumClosest(nums[1:], nums[0])
		if res != answer {
			t.Fatalf("%v excepted %d, but got %d", nums[1:], answer, res)
		}
	}
}
