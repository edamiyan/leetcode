package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp, dummy := 0, new(ListNode)
	for node := dummy; l1 != nil || l2 != nil; node = node.Next {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{Val: tmp % 10}
		tmp /= 10
	}
	return dummy.Next
}
