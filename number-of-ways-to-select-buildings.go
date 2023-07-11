//https://leetcode.com/problems/number-of-ways-to-select-buildings/

package main

import "fmt"

// the only possible combinations we need to find "101" and "010"
// every combinatiщn we need to research based on the current number
// if current number is 0 then we need to find how many 1 before and after this - this will be "101" case
// if current number is 1 then we need to find how many 0 before and after this - this will be "010" case
func numberOfWays(s string) int64 {

	result := 0

	tZeroes := 0

	//calculate how many zeroes at total
	for _, ch := range s {
		if ch == '0' {
			tZeroes++
		}
	}

	tOnes := len(s) - tZeroes

	// curr values show number of zeroes or ones before the current value
	cZeroes := 0
	cOnes := 0

	for i, ch := range s {
		if i == 0 {
			if ch == '0' {
				cZeroes++
			} else {
				cOnes++
			}
			continue
		}

		if ch == '0' { //101 case
			result += cOnes * (tOnes - cOnes)
			cZeroes++
		} else { //010 case
			result += cZeroes * (tZeroes - cZeroes)
			cOnes++
		}
	}

	return int64(result)
}

func main() {

	fmt.Printf("%d should be 6", numberOfWays("0011011"))
}
