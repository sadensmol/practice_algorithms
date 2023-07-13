// https://www.codewars.com/kata/59c633e7dcc4053512000073/train/go
package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func solve(str string) int {
	maxSum := 0
	curSum := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			if curSum > maxSum {
				maxSum = curSum
			}
			curSum = 0
		} else {
			curSum += int(str[i]) - 96
		}
	}

	if curSum > maxSum {
		maxSum = curSum
	}

	return maxSum
}

var _ = Describe("Test Example", func() {
	It("Should return the correct values!", func() {
		Expect(solve("a")).To(Equal(0))
		Expect(solve("bcd")).To(Equal(9))
		Expect(solve("zea")).To(Equal(26))
		Expect(solve("az")).To(Equal(26))
		Expect(solve("baz")).To(Equal(26))
		Expect(solve("aeiou")).To(Equal(0))
		Expect(solve("uaoczei")).To(Equal(29))
		Expect(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")).To(Equal(143))
		Expect(solve("codewars")).To(Equal(37))
		Expect(solve("bup")).To(Equal(16))
	})
})

func TestConsonantValue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Example")
}
