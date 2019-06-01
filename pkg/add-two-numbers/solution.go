package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cursor := res
	tail := 0 // 进位
	cursorA, cursorB := l1, l2
	for {
		tmp := 0
		if cursorB != nil {
			tmp += cursorB.Val
			cursorB = cursorB.Next
		}

		if cursorA != nil {
			tmp += cursorA.Val
			cursorA = cursorA.Next
		}

		cursor.Val = (tmp + tail) % 10
		tail = (tmp + tail) / 10
		if cursorA != nil || cursorB != nil || tail != 0 {
			cursor.Next = &ListNode{}
		} else {
			break
		}
		cursor = cursor.Next
	}

	return res
}
