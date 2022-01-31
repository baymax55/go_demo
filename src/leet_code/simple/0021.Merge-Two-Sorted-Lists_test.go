package leetcode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1_1 := ListNode{Val: 2}
	l1 := ListNode{1, &l1_1}

	l4_1 := ListNode{Val: 4}
	l4 := ListNode{3, &l4_1}
	res := mergeTwoLists(&l4, &l1)
	PL(res)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
