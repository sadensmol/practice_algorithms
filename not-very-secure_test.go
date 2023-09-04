// https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go

package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, c := range str {
		if (c < '0' || c > '9') && (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return false
		}
	}

	return true
}

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(alphanumeric(".*?")).To(Equal(false))
		Expect(alphanumeric("a")).To(Equal(true))
		Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
		Expect(alphanumeric("hello world_")).To(Equal(false))
		Expect(alphanumeric("PassW0rd")).To(Equal(true))
		Expect(alphanumeric("     ")).To(Equal(false))
		Expect(alphanumeric("")).To(Equal(false))
		Expect(alphanumeric("\n\t\n")).To(Equal(false))
		Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
		Expect(alphanumeric("__ * __")).To(Equal(false))
		Expect(alphanumeric("&)))(((")).To(Equal(false))
		Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
	})
})

func TestNotVerySecure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test")
}
