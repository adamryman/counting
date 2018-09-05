package counting

import "testing"
import "math/rand"
import "sort"

func TestSort(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		max         int
		expected    []int
	}{
		{
			"Simple",
			[]int{1, 4, 1, 2, 7, 5, 2},
			9,
			[]int{1, 1, 2, 2, 4, 5, 7},
		},
	}

	for _, test := range tests {
		actual := Sort(test.input, test.max)

		for i := 0; i < len(actual); i++ {
			if actual[i] != test.expected[i] {
				t.Log(actual)
				t.Log(test.expected)
				t.Fatal()
			}
		}
	}
}

func RandomIntSlice(maxSize, maxLength int) []int {
	out := make([]int, maxLength, maxLength)
	for i, _ := range out {
		out[i] = rand.Intn(maxSize)
	}
	return out
}

func benchmarkSort(maxSize, maxLength int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(RandomIntSlice(maxSize, maxLength), maxSize)
	}
}

func benchmarkSortStdlib(maxSize, maxLength int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort.Ints(RandomIntSlice(maxSize, maxLength))
	}
}

func BenchmarkSortStdlib1(b *testing.B) { benchmarkSortStdlib(1, 1, b) }
func BenchmarkSort1(b *testing.B)       { benchmarkSort(1, 1, b) }

func BenchmarkSortStdlib2(b *testing.B) { benchmarkSortStdlib(2, 2, b) }
func BenchmarkSort2(b *testing.B)       { benchmarkSort(2, 2, b) }

func BenchmarkSortStdlib3(b *testing.B) { benchmarkSortStdlib(3, 3, b) }
func BenchmarkSort3(b *testing.B)       { benchmarkSort(3, 3, b) }

func BenchmarkSortStdlib10(b *testing.B) { benchmarkSortStdlib(10, 10, b) }
func BenchmarkSort10(b *testing.B)       { benchmarkSort(10, 10, b) }

func BenchmarkSortStdlib20(b *testing.B) { benchmarkSortStdlib(20, 20, b) }
func BenchmarkSort20(b *testing.B)       { benchmarkSort(20, 20, b) }

func BenchmarkSortStdlib40(b *testing.B) { benchmarkSortStdlib(40, 40, b) }
func BenchmarkSort40(b *testing.B)       { benchmarkSort(40, 40, b) }

func BenchmarkSort10And1000Stdlib(b *testing.B) { benchmarkSortStdlib(10, 1000, b) }
func BenchmarkSort10And1000(b *testing.B)       { benchmarkSort(10, 1000, b) }

func BenchmarkSort10And10000Stdlib(b *testing.B) { benchmarkSortStdlib(10, 10000, b) }
func BenchmarkSort10And10000(b *testing.B)       { benchmarkSort(10, 10000, b) }

func BenchmarkSort10And100000Stdlib(b *testing.B) { benchmarkSortStdlib(10, 100000, b) }
func BenchmarkSort10And100000(b *testing.B)       { benchmarkSort(10, 100000, b) }

func BenchmarkSort10And1000000Stdlib(b *testing.B) { benchmarkSortStdlib(10, 1000000, b) }
func BenchmarkSort10And1000000(b *testing.B)       { benchmarkSort(10, 1000000, b) }

func BenchmarkSort100And1000000Stdlib(b *testing.B) { benchmarkSortStdlib(100, 1000000, b) }
func BenchmarkSort100And1000000(b *testing.B)       { benchmarkSort(100, 1000000, b) }

func BenchmarkSort1000And1000000Stdlib(b *testing.B) { benchmarkSortStdlib(1000, 1000000, b) }
func BenchmarkSort1000And1000000(b *testing.B)       { benchmarkSort(1000, 1000000, b) }
