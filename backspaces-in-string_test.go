package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func CleanString(s string) string {
	r := make([]byte, len(s))
	p := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if p > 0 {
				p--
			}
		} else {
			r[p] = s[i]
			p++
		}
	}

	if p <= 0 {
		return ""
	}

	return string(r[:p])
}

var _ = Describe("Test", func() {
	It("Sample tests", func() {
		Expect(CleanString("abc#d##c")).To(Equal("ac"))
		Expect(CleanString("abc####d##c#")).To(Equal(""))
		Expect(CleanString("")).To(Equal(""))
		Expect(CleanString("#######")).To(Equal(""))
	})
})

func TestBackspacesInString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test")
}
