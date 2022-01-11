package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1_1 := ListNode{Val: 1}
	l2_1 := ListNode{Val: 1}

	l1 := ListNode{1, &l1_1}
	l2 := ListNode{2, &l2_1}
	l := addTwoNumbers(&l1, &l2)
	p(l)
}
func p(l *ListNode) {
	inner := l.Next
	val := l.Val
	println(val)
	if inner == nil {
		return
	}
	p(inner)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
