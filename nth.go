package nth

import "sort"

// Element finds the nth rank ordered element and ensures it is at the nth
// position.
func Element(data sort.Interface, n int) {
	if n < 0 || n >= data.Len() {
		return
	}
	sort.Sort(data)
}
