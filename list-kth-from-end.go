package main

import "fmt"

type Item struct {
	val  int
	next *Item
}

func main() {
	// 50 98 51 71 19 4
	head := &Item{50, &Item{98, &Item{51, &Item{95, &Item{71, &Item{19, &Item{4, nil}}}}}}}

	fmt.Println(kthFromEnd(0, head)) //4
	fmt.Println(kthFromEnd(1, head)) // 19
	fmt.Println(kthFromEnd(3, head)) //95
	fmt.Println(kthFromEnd(6, head)) //50
	fmt.Println(kthFromEnd(9, head)) //nil
}

func kthFromEnd(k int, head *Item) *Item {
	kp := head

	for i := 0; i < k; i++ {
		if kp.next == nil {
			return nil
		}
		kp = kp.next
	}

	p := head
	for {
		if kp.next == nil {
			return p
		}
		kp = kp.next
		p = p.next
	}

	return p
}
