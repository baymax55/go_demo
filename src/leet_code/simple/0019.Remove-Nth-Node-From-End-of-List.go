package leetcode

import (
	structures "leet_code/self_pkg"
)

func PL(l *ListNode) {
	inner := l.Next
	val := l.Val
	print(val, " ")
	if inner == nil {
		return
	}
	PL(inner)
}

//func main() {
//
//	l1 := ListNode{Val: 1}
//	l2 := ListNode{2, &l1}
//	l3 := ListNode{3, &l2}
//	l4 := ListNode{4, &l3}
//	PL(&l4)
//	println()
//	res := removeNthFromEnd(&l4, 1)
//	PL(res)
//}

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}
