package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type List struct {
	start *ListNode
}

func addTwoNumberss(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}

		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var res *ListNode
	for l1 != nil || l2 != nil {

		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		res.Val = sum % 10
		res = res.Next
	}

	return res
}

/*

absolute hire

no hire but you can convince me

*/
