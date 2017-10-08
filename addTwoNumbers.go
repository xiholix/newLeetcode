package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) show() {
	ptr := this

	for ptr != nil {
		fmt.Printf("%d -> ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	result := ListNode{}
	tmp := &result
	var endPtr *ListNode

	for l1 != nil || l2 != nil {
		var v1 int = 0
		var v2 int = 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		v := v1 + v2 + carry
		tmp.Val = v % 10

		carry = v / 10

		t := ListNode{}
		tmp.Next = &t
		endPtr = tmp
		tmp = &t

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		endPtr.Next.Val = carry
	} else {
		endPtr.Next = nil
	}

	return &result
}

func main() {
	a := ListNode{Val: 3, Next: nil}
	b := ListNode{Val: 4, Next: &a}
	c := ListNode{Val: 2, Next: &b}

	a1 := ListNode{Val: 4, Next: nil}
	b1 := ListNode{Val: 6, Next: &a1}
	c1 := ListNode{Val: 5, Next: &b1}

	c1.show()
	c.show()

	t := addTwoNumbers(&c1, &c)
	t.show()
}
