package sortreverse

import (
	"fmt"
	"math/rand"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func BenchmarkSortReverse1(b *testing.B) {
	var _ = func() bool {
		testing.Init()
		return true
	}()

	a := make([]int, 100_000)
	for n := 0; n < b.N; n++ {
		for i := range a {
			a[i] = rand.Intn(100_000)
		}
		sortReverse1(a)
	}
}

func BenchmarkSortReverse2(b *testing.B) {
	var _ = func() bool {
		testing.Init()
		return true
	}()

	a := make([]int, 100_000)
	for n := 0; n < b.N; n++ {
		for i := range a {
			a[i] = rand.Intn(100_000)
		}
		sortReverse2(a)
	}
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
