package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	val  int
	next *Item
}

func (it *Item) String() (result string) {
	for cur := it; cur != nil; cur = cur.next {
		result += strconv.Itoa(cur.val)
	}
	return result
}

func main() {

	//2->4->3 = 242
	first := &Item{2, &Item{4, &Item{3, nil}}}
	// 5->6->4 = 465
	second := &Item{5, &Item{6, &Item{4, nil}}}
	zero := &Item{0, nil}
	// 9->9->9->9->9->9->9 = 9999999
	third := &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, nil}}}}}}}
	// 9->9->9->9 = 9999
	fourth := &Item{9, &Item{9, &Item{9, &Item{9, nil}}}}

	fmt.Println(addTwoNumbers(first, second)) // 7->0->8 = 807
	fmt.Println(addTwoNumbers(zero, zero))    // 0
	fmt.Println(addTwoNumbers(third, fourth)) // 8->9->9->9->0->0->0->1 = 10009998
}

func addTwoNumbers(a, b *Item) *Item {
	var result, prev *Item
	carry := 0

	for a != nil || b != nil {
		sum := carry
		if a != nil {
			sum += a.val
			a = a.next
		}
		if b != nil {
			sum += b.val
			b = b.next
		}
		carry = sum / 10
		sum = sum % 10

		if result == nil {
			result = &Item{sum, nil}
			prev = result
		} else {
			prev.next = &Item{sum, nil}
			prev = prev.next
		}
	}

	if carry > 0 {
		prev.next = &Item{carry, nil}
	}

	return result
}
