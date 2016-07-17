package nth

import (
	"sort"
	"testing"
)

var shuffled = []int{10, 14, 6, 7, 16, 12, 9, 0, 8, 4, 11, 5, 15, 1, 2, 13, 3}
var asc = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
var desc = []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func TestElement(t *testing.T) {
	cases := map[string][]int{
		"shuffled": shuffled,
		"asc":      asc,
		"desc":     desc,
	}
	for name, src := range cases {
		data := make([]int, len(src))
		for n := range src {
			copy(data, src)
			Element(sort.IntSlice(data), n)
			if data[n] != n {
				t.Errorf("%s: Element(..., %d) != %d: %v", name, n, n, data)
			}
		}
	}
}

func BenchmarkElement(b *testing.B) {
	data := make([]int, len(shuffled))
	dataS := sort.IntSlice(data)
	for n := 0; n < b.N; n++ {
		copy(data, shuffled)
		Element(dataS, 17)
	}
}
