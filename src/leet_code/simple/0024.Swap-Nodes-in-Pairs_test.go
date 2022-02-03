package leetcode

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	l2 := ListNode{Val: 2}
	l1 := ListNode{1, &l2}

	l3 := ListNode{Val: 3}
	l4 := ListNode{4, &l3}
	l3.Next = &l1
	PL(&l4)
	println()
	PL(swapPairs(&l4))
}
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
