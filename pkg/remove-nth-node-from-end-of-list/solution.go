package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cursor := 0
	nodeCursor, nthNode := head, head
	for nodeCursor != nil {
		// 这里维持第二个指针，永远比最后一个指针靠前 n 个节点
		if cursor > n {
			nthNode = nthNode.Next
		}
		nodeCursor = nodeCursor.Next
		cursor++
	}

	// 刚好移除第一个节点
	if cursor == n {
		return head.Next
	}

	nthNode.Next = nthNode.Next.Next
	return head
}
