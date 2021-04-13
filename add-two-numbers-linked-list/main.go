package main

import "fmt"

type ListNode struct {
	Val  int64
	Next *ListNode
}
type List struct {
	start *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var tmp1 int64
	var mul int64
	mul = 1
	for l1 != nil {
		a := int64(l1.Val)
		a = a * mul
		tmp1 = a + tmp1
		mul *= 10
		l1 = l1.Next
	}
	fmt.Println("tmp1", tmp1)
	var tmp2 int64
	mul = 1
	for l2 != nil {
		a := int64(l2.Val)
		a = a * mul
		tmp2 = a + tmp2
		mul *= 10
		l2 = l2.Next
	}

	fmt.Println("tmp2", tmp2)
	total := tmp1 + tmp2
	//

	fmt.Println(total)
	var tmp3 int64
	var newVal int64
	l3 := &List{}

	if total == 0 {
		var node ListNode
		node.Val = 0
		l3.start = &node
		return l3.start
	}

	for total > 0 {
		remainder := int64(total % 10)
		newVal = (remainder + tmp3)
		tmp3 = newVal * 10

		var node ListNode
		//fmt.Println("remainder",remainder)
		node.Val = remainder

		if l3.start == nil {
			fmt.Println("hello")
			l3.start = &node
		} else {
			fmt.Println("remainder", remainder)
			list := l3.start
			for list.Next != nil {
				list = list.Next
			}
			list.Next = &node
		}

		total /= 10
	}

	fmt.Println(newVal)
	// fmt.Println(l3)
	//l3 := ListNode{}

	// for l3 != nil {
	//     fmt.Println("l3.Val",l3.Val)
	//     l3 = l3.Next
	// }

	return l3.start
}

// func Add(l1 *ListNode) *ListNode {

//     if l1 != nil {
//         l1.
//     }
// }
