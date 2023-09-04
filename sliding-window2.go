package main

// You are given a binary string s (a string containing only "0" and "1"). You may choose up to one "0" and flip it to a "1".
// What is the length of the longest substring achievable that contains only "1"?
// For example, given s = "1101100111", the answer is 5. If you perform the flip at index 2, the string becomes 1111100111.

func main() {
	s := "1101100111"
	println("Longest substring length: ", longestSubstring(s))
}

func longestSubstring(s string) int {
	start, end, maxLen, count := 0, 0, 0, 0
	for end < len(s) {
		if s[end] == '0' {
			count++
		}
		for count > 1 {
			if s[start] == '0' {
				count--
			}
			start++
		}
		maxLen = max(maxLen, end-start+1)
		end++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
