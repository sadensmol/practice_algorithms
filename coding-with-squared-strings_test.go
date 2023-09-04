// https://www.codewars.com/kata/56fcc393c5957c666900024d
// todo

package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Code(s string) string {
	l := len(s)

	i := 1 // smallest square
	for {
		if i*i >= l {
			break
		}
		i++
	}

	return ""
}
func Decode(s string) string {
	// your code
	return ""
}

func dotestCode(a1 string, exp string) {
	var ans = Code(a1)
	Expect(ans).To(Equal(exp))
}
func dotestDecode(a1 string, exp string) {
	var ans = Decode(a1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Code, Decode", func() {

	It("should handle basic cases", func() {
		var d = "I.was.going.fishing.that.morning.at.ten.o'clock"
		var s = "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"
		dotestCode(d, s)
		dotestDecode(s, d)

		d = "Process terminated with status 0 (0 minute(s), 6 second(s))"
		s = "s t setP\n)se(tder\n)e(0a ro\n\vcs twmc\n\vo)muiie\n\vn,istns\n\vd n has\n\v(6u0 t "
		dotestCode(d, s)
		dotestDecode(s, d)

	})
})

// add test
func TestCodingWithSquaredStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Example")
}
