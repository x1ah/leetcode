package swap_nodes_in_pairs

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	nodes := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	res := swapPairs(nodes)
	answers := []int{2, 1, 4, 3}
	idx := 0
	for res != nil {
		if res.Val != answers[idx] {
			t.Fatalf("%v excepted %v, but got %v", nodes, answers, res)
		}
		res = res.Next
		idx++
	}
}
