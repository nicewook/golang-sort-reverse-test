package sortreverse

import (
	"sort"
)

func sortReverse1(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

func sortReverse2(a []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
}
