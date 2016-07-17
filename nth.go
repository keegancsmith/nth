package nth

import "sort"

// Element finds the nth rank ordered element and ensures it is at the nth
// position.
func Element(data sort.Interface, n int) {
	l := data.Len()
	if n < 0 || n >= l {
		return
	}
	quickSelect(data, n, 0, l)
}

func quickSelect(data sort.Interface, n, a, b int) {
	pivot := a
	for i := a + 1; i < b; i++ {
		if data.Less(i, pivot) { // data[i] < data[pivot]
			// data[i] needs to be before data[pivot]
			data.Swap(pivot, i)
			pivot++
			data.Swap(pivot, i)
		}
	}
	if a+n < pivot {
		quickSelect(data, n, a, pivot)
	} else if pivot < a+n {
		quickSelect(data, a+n-pivot-1, pivot+1, b)
	}
}
