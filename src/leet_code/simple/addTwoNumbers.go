package leetcode

/**
 * Definition for singly-linked list.
 * type listNode struct {
 *     Val int
 *     Next *listNode
 * }
 */

type listNode struct {
	Val  int
	Next *listNode
}

//func main() {
//	l1_1 := listNode{Val: 9}
//	l2_1 := listNode{Val: 1}
//
//	l1 := listNode{9, &l1_1}
//	l2 := listNode{1, &l2_1}
//	l := addTwoNumbers(&l1, &l2)
//	p(l)
//}
func p(l *listNode) {
	inner := l.Next
	val := l.Val
	println(val)
	if inner == nil {
		return
	}
	p(inner)
}
func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	head := &listNode{Val: 0}
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
		current.Next = &listNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
