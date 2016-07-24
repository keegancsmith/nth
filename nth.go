package nth

import "sort"

// Element finds the nth rank ordered element and ensures it is at the nth
// position.
func Element(data sort.Interface, n int) {
	l := data.Len()
	if n < 0 || n >= l {
		return
	}
	quickSelectAdaptive(data, n, 0, l)
}

// quickSelectAdaptive is from "Fast Deterministic Selection" by Andrei
// Alexandrescu https://arxiv.org/abs/1606.00484
//
// It is deterministic O(n) selection algorithm tuned for real world
// workloads.
func quickSelectAdaptive(data sort.Interface, k, a, b int) {
	var (
		l int // |A| from the paper
		p int // pivot position
	)
	for {
		l = b - a
		r := float64(k) / float64(l) // r <- real(k) / real(|A|)
		if l < 12 {
			p = hoarePartition(data, a+l/2, a, b) // HoarePartition(A, |A| / 2)
		} else if r < 7.0/16.0 {
			if r < 1.0/12.0 {
				p = repeatedStepFarLeft(data, k, a, b)
			} else {
				p = repeatedStepLeft(data, k, a, b)
			}
		} else if r >= 1.0-7.0/16.0 {
			if r >= 1.0-1.0/12.0 {
				p = repeatedStepFarRight(data, k, a, b)
			} else {
				p = repeatedStepRight(data, k, a, b)
			}
		} else {
			p = repeatedStepImproved(data, k, a, b)
		}
		if p == k {
			return
		}
		if p > k {
			b = p // A <- A[0:p]
		} else {
			// i <- k - p - 1  // TODO what is i?
			a = p + 1 // A <- A[p+1:|A|]
		}
	}
}

func hoarePartition(data sort.Interface, p, begin, end int) int {
	data.Swap(p, begin) // Swap(A[p], A[0])
	a := begin + 1      // a = 1
	b := end - 1        // b = |A| - 1
Loop:
	for {
		for {
			if a > b {
				break Loop
			}
			if !data.Less(a, begin) { // A[a] >= A[0]
				break
			}
			a++
		}
		for data.Less(begin, b) { // A[0] < A[b]
			b--
		}
		if a >= b {
			break
		}
		data.Swap(a, b) // Swap(A[a], A[b])
		a++
		b--
	}
	data.Swap(begin, a-1) // Swap(A[0], A[a-1])
	return a - 1
}

func simplePartition(data sort.Interface, k, a, b int) int {
	p := a
	for i := a + 1; i < b; i++ {
		if data.Less(i, p) {
			data.Swap(p, i)
			p++
			data.Swap(p, i)
		}
	}
	return p
}

// TODO implement
var repeatedStepFarLeft = simplePartition
var repeatedStepLeft = simplePartition
var repeatedStepFarRight = simplePartition
var repeatedStepRight = simplePartition
var repeatedStepImproved = simplePartition
