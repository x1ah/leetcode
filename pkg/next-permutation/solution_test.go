package next_permutation

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	input := []int{5, 1, 1}
	answer := []int{1, 1, 5}
	nextPermutation(input)
	for i, v := range input {
		if v != answer[i] {
			t.Fatalf("excepted %v, got %v", answer, input)
		}
	}
}
