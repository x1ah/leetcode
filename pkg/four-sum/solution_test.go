package four_sum

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	k := fourSum(input, 0) // excepted [[-1 1 2 -2] [0 0 2 -2] [0 0 1 -1]]
	fmt.Println(k)
}
