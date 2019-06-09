package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	l := removeElement(input, 2)
	if l != 5 {
		t.Fatalf("length excepted 5, got %d", l)
	}
}
