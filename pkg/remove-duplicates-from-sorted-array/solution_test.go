package remove_duplicates_from_sorted_array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	inputs := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(inputs)
	if length != 5 {
		t.Fatalf("length excepted 5, got %d", length)
	}

	answer := []int{0, 1, 2, 3, 4}
	for i, v := range answer {
		if v != inputs[i] {
			t.Fatalf("excepted %v, bug got %v", answer, inputs[:5])
		}
	}
}
