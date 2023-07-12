// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	r := PosPeaks{}
	if len(array) < 3 {
		return r
	}

	prev := array[0]
	f := false // we are growing
	pI := 0    // platoe position
	for i, v := range array {
		if v > prev { //we are growing
			pI = 0
			f = true
		} else if v == prev {
			if pI == 0 {
				pI = i - 1
			}
		} else { // we are falling
			if f == false { // // still falling

			} else {
				if pI > 0 { // was plateau
					r.Pos = append(r.Pos, pI)
				} else {
					r.Pos = append(r.Pos, i-1)
				}
				pI = 0
				r.Peaks = append(r.Peaks, prev)
				f = false
			}
		}
		prev = v
	}
	return r
}

func dotest(array []int, expected PosPeaks) {
	var ans = PickPeaks(array)
	if len(expected.Pos) == 0 {
		if len(ans.Pos) != 0 || len(ans.Peaks) != 0 {
			Expect(ans).To(Equal(expected))
		} else {
			Expect(true).To(BeTrue())
		}
	} else {
		Expect(ans).To(Equal(expected))
	}
}

var _ = Describe("Test Pick Peaks", func() {
	It("should support finding peaks", func() {
		dotest(
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})
})

func TestPickPeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Pick Peaks")
}
