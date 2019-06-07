package merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	b := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	// 1, 1, 2, 3, 4, 4
	res := mergeTwoLists(a, b)
	for r := res; r != nil; r = r.Next {
		fmt.Println(r.Val)
	}
}
