package main

import "fmt"

// Given an array of positive integers nums and an integer k,
// find the length of the longest subarray whose sum is less than or equal to k.

func main() {
	a := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	k := 8
	fmt.Println("Longest subarray length: ", longestSubarray(a, k))
}

func longestSubarray(nums []int, k int) int {
	start, end, sum, maxLen := 0, 0, 0, 0
	for end < len(nums) {
		sum += nums[end]
		for sum > k {
			sum -= nums[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
		end++
	}
	fmt.Printf("from index: %d to index %d\n", start-1, end-1)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
