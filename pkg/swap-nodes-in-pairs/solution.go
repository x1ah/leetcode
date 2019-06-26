package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	var res *ListNode
	if head != nil && head.Next != nil {
		res = head.Next
	} else {
		return head
	}
	cursor := head
	var pre *ListNode
	for cursor != nil && cursor.Next != nil {
		tmp := cursor.Next
		cursor.Next = tmp.Next
		tmp.Next = cursor
		if pre != nil {
			pre.Next = tmp
		}
		pre = cursor
		cursor = cursor.Next
	}

	return res
}
