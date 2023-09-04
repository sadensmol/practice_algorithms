package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	a := []int{11, 1, 22, 2.0, 5, 6, 5, 9, 7, 10, 3}

	slices.Sort(a)
	fmt.Println(a)
}
