package sortreverse

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSortReverse(b *testing.B) {
	a := make([]int, 100_000)
	for i := range a {
		a[i] = rand.Intn(100_000)
	}

	b.Run("sort using sort.Reverse", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sortReverse1(a)
		}
	})
	b.Run("sort using sort.Slice", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sortReverse2(a)
		}
	})
}

func ExampleSort() {
	a := []int{1, 2, 3, 4, 5}
	sortReverse1(a)
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5}
	sortReverse2(a)
	fmt.Println(a)
	// Output:
	// [5 4 3 2 1]
	// [5 4 3 2 1]
}
