package main

import "fmt"

// Дан отсортированный массив чисел, включая отрицательные, и некое число K.
// Нужно найти два числа из массива, которые в сумме дадут число K.
// Если таких чисел нет, в результате возвращаем пустой массив.
// Если комбинаций чисел, дающих нужную сумму, несколько, возвращаем любую из них.

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var k = 13

	var result = sum2Digits(arr, k)
	fmt.Println(result)

}

// hash based solution
func sum2Digits(arr []int, k int) []int {
	m := make(map[int]bool)

	for _, v := range arr {
		r := k - v // we need this number to get k

		if _, ok := m[r]; ok {
			return []int{r, v}
		}
		m[v] = true
	}
	return []int{}
}
