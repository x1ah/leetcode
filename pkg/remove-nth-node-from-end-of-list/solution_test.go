package remove_nth_node_from_end_of_list

import "testing"

func TestRemoveNthFormEnd(t *testing.T) {
	testcases := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		{Val: 1, Next: nil},
	}
	n := []int{2, 1}
	//answers := []*ListNode{
	//	&ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:5, Next:nil}}}},
	//	nil,
	//}

	for i := 0; i < 2; i++ {
		res := removeNthFromEnd(testcases[i], n[i])
		if i == 1 {
			if res != nil {
				t.Fatalf("%v excepted nil", testcases[i])
			}
		}
	}
}
