// https://www.codewars.com/kata/559536379512a64472000053/train/go
package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func PlayPass(s string, n int) string {
	r := make([]byte, len(s))
	j := 0
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] >= 'A' && s[i] <= 'Z' {
			v := s[i]
			if i%2 == 1 {
				v = v + 32
				r[j] = byte('a' + (v-'a'+byte(n))%26)
			} else {
				r[j] = byte('A' + (v-'A'+byte(n))%26)
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			v := s[i]
			if i%2 == 0 {
				v = v - 32
				r[j] = byte('A' + (v-'A'+byte(n))%26)
			} else {
				r[j] = byte('a' + (v-'a'+byte(n))%26)
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			r[j] = byte('0' + '9' - s[i])
		} else {
			r[j] = s[i]
		}

		j++
	}

	return string(r)
}

func dotest2(a1 string, n int, exp string) {
	var ans = PlayPass(a1, n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests PlayPass", func() {
	It("simple text", func() {
		dotest2("abcd", 0, "dCbA")
		dotest2("abc", 1, "DcB")
	})

	It("should handle basic cases", func() {
		dotest2("I LOVE YOU!!!", 1, "!!!vPz fWpM J")
		dotest2("I LOVE YOU!!!", 0, "!!!uOy eVoL I")
		dotest2("AAABBCCY", 1, "zDdCcBbB")

	})
})

func TestPlayingWithPassphrases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test PlayPass")
}
