package leetcode

import "testing"

func TestReverseKGroup(t *testing.T) {
	l2 := ListNode{Val: 2}
	l1 := ListNode{1, &l2}

	l3 := ListNode{Val: 3}
	l4 := ListNode{4, &l3}
	l3.Next = &l1
	PL(&l4)
	println()
	PL(reverseKGroup(&l4, 2))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
