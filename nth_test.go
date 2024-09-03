package nth

import (
	"fmt"
	"sort"
	"testing"
	"testing/quick"
)

var cases = []struct {
	name string
	data []byte
}{{
	name: "shuffled",
	data: []byte{10, 14, 6, 7, 16, 12, 9, 0, 8, 4, 11, 5, 15, 1, 2, 13, 3},
}, {
	name: "asc",
	data: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
}, {
	name: "desc",
	data: []byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}, {
	name: "fuzz20240831",
	data: []byte{48, 48, 48, 48, 48, 48, 48, 32, 48, 32, 32, 48, 48, 48, 48, 32, 48, 48, 33, 48, 48, 48, 48, 48, 48, 32, 32, 32, 32},
}}

func TestElement(t *testing.T) {
	for _, tc := range cases {
		src := append([]byte{}, tc.data...)
		data := make([]byte, len(src))
		for n := range src {
			copy(data, src)
			Element(byteSlice(data), n)
			if data[n] != byte(n) {
				t.Errorf("%s: Element(..., %d) != %d: %v", tc.name, n, n, data)
			}
		}
	}
}

func TestElementQuick(t *testing.T) {
	f := func(data []byte, n uint) bool {
		return checkElement(t, data, n)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func FuzzElement(f *testing.F) {
	for _, tc := range cases {
		f.Add(tc.data, uint(9))
	}
	f.Fuzz(func(t *testing.T, data []byte, n uint) {
		if !checkElement(t, data, n) {
			t.Error()
		}
	})
}

func checkElement(t interface{ Logf(string, ...any) }, data []byte, n uint) bool {
	if len(data) == 0 {
		return true
	}
	n = n % uint(len(data)) // Ensure n is within the bounds of the slice

	got := append([]byte{}, data...)
	Element(byteSlice(got), int(n))

	sorted := append([]byte{}, data...)
	sort.Sort(byteSlice(sorted))

	if got[n] == sorted[n] {
		return true
	}

	t.Logf("Element(%v, %d) returned an incorrect answer", data, n)
	t.Logf("got:    %v", got)
	t.Logf("sorted: %v", sorted)
	t.Logf("(got[%d] = %d) != (sorted[%d] = %d)", n, got[n], n, sorted[n])
	return false
}

func BenchmarkElement(b *testing.B) {
	shuffled := append([]byte{}, cases[0].data...)
	data := make([]byte, len(shuffled))
	dataS := byteSlice(data)
	for n := 0; n < b.N; n++ {
		copy(data, shuffled)
		Element(dataS, 15)
	}
}

func TestPartition(t *testing.T) {
	testPartition(t, "hoarePartition", hoarePartition)
	testPartition(t, "simplePartition", simplePartition)
	testPartition(t, "repeatedStepFarLeft", repeatedStepFarLeft)
	testPartition(t, "repeatedStepLeft", repeatedStepLeft)
}

func testPartition(t *testing.T, name string, f func(sort.Interface, int, int, int) int) {
	t.Run(name, func(t *testing.T) {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				src := append([]byte{}, tc.data...)
				data := make([]byte, len(src))
				for a := 0; a < len(src); a++ {
					for b := a + 1; b <= len(src); b++ {
						for k := a; k < b; k++ {
							copy(data, src)
							want := data[k]
							p := f(byteSlice(data), k, a, b)
							got := data[p]
							cname := fmt.Sprintf("(..., %d, %d, %d) = %d", k, a, b, p)
							if p < a || p >= b {
								t.Errorf("%s not in range [a,b)", cname)
							}
							for i := a; i < p; i++ {
								if data[i] > data[p] {
									t.Errorf("%s not partitioned. A[%d] > A[%d] = A[p]", cname, i, p)
								}
							}
							for i := p + 1; i < b; i++ {
								if data[i] < data[p] {
									t.Errorf("%s not partitioned. A[%d] < A[%d] = A[p]", cname, i, p)
								}
							}
							if want != got {
								t.Errorf("%s did not partion around k. got %v != %v", cname, got, want)
							}
						}
					}
				}
			})
		}
	})
}

type byteSlice []byte

func (x byteSlice) Len() int           { return len(x) }
func (x byteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x byteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
