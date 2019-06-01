package add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testcases := [][3]*ListNode{
		{
			&ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			&ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}, // answer
		},
		{
			&ListNode{Val: 5, Next: nil},
			&ListNode{Val: 5, Next: nil},
			&ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}},
		},
	}

	for _, testcase := range testcases {
		res := addTwoNumbers(testcase[0], testcase[1])
		cursorA, cursorB := testcase[2], res
		for times := 0; times < 100; times++ {
			if cursorA == nil && cursorB == nil {
				break
			}

			if cursorB.Val != cursorA.Val {
				t.Fatalf("%v fail, got %v", testcase, res)
			}

			cursorA = cursorA.Next
			cursorB = cursorB.Next
		}
	}
}
