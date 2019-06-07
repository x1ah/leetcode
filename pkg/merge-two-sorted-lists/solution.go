package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	cursorA, cursorB, cursorC := l1, l2, newNode
	for cursorB != nil && cursorA != nil {
		if cursorB.Val <= cursorA.Val {
			cursorC.Next = cursorB
			cursorB = cursorB.Next
		} else {
			cursorC.Next = cursorA
			cursorA = cursorA.Next
		}
		cursorC = cursorC.Next
	}

	if cursorA != nil {
		cursorC.Next = cursorA
	}

	if cursorB != nil {
		cursorC.Next = cursorB
	}

	return newNode.Next
}
