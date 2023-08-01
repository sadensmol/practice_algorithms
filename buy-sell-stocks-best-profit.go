package main

import "math"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func maxProfit(prices []int) int {
	m := math.MaxInt64 // min
	p := 0             // profit

	for i := 0; i < len(prices); i++ {
		c := prices[i]
		if c < m {
			m = prices[i]
		}

		cp := c - m
		if cp > p {
			p = cp
		}

	}
	return p
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
