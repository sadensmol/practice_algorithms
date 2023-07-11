package main

import "fmt"

func main() {
	println("test")
	v4 := ListNode{Val: 4}
	v3 := ListNode{Val: 3, Next: &v4}
	v2 := ListNode{Val: 2, Next: &v3}
	v1 := ListNode{Val: 1, Next: &v2}
	ttt := swapPairs(&v1)
	fmt.Printf("%#v", ttt)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	ttt := head.Next
	vvv := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairs(vvv)
	return ttt

}
