package leetcode

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {

	l1_1 := ListNode{Val: 2}
	l1 := ListNode{1, &l1_1}

	l4_1 := ListNode{Val: 4}
	l4 := ListNode{3, &l4_1}

	l5 := ListNode{Val: 6}

	l := []*ListNode{&l1, &l4, &l5}

	srr := mergeKLists(l)
	PL(srr)
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists1(left, right)
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}
