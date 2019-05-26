package three_sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	inputs := [][]int{
		{1, 1, -2},
		{-1, 0, 1, 2, -1, -4},
		{-2, 0, 0, 2, 2},
		{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
		{0, 0, 0},
	}

	outputs := [][][]int{
		{
			{-2, 1, 1},
		},
		{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		{
			{-2, 0, 2},
		},
		{
			{-5, 1, 4},
			{-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0},
		},
		{
			{0, 0, 0},
		},
	}

	for idx, input := range inputs {
		res := threeSum(input)
		fmt.Println(res, outputs[idx])
	}
}
