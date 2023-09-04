// На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение

package main

import "fmt"

func main() {
	a := []int{11, 3, 6, 4, 5, 7, 8, 9, 10}
	b := []int{1, 22, 11, 4, 123, 66, 5, 7, 8, 9, 10}

	fmt.Println(cross(a, b))
	fmt.Println(crossImproved(a, b))
}
func cross(a []int, b []int) []int {
	var r []int
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				r = append(r, v)
			}
		}
	}
	return r
}

func crossImproved(a []int, b []int) []int {
	var r []int
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			r = append(r, v)
		}
	}
	return r
}
