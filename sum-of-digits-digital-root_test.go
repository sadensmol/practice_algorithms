// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	v := n % 10
	p := (n - v) / 10

	return DigitalRoot(v + DigitalRoot(p))

}

var _ = Describe("Test Example", func() {
	It("fixed tests", func() {
		Expect(DigitalRoot(16)).To(Equal(7))
		Expect(DigitalRoot(195)).To(Equal(6))
		Expect(DigitalRoot(992)).To(Equal(2))
		Expect(DigitalRoot(167346)).To(Equal(9))
		Expect(DigitalRoot(0)).To(Equal(0))
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Example")
}
