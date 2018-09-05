# Counting Sort [![Build Status](https://travis-ci.org/adamryman/counting.svg?branch=master)](https://travis-ci.org/adamryman/counting)

This is an implementation of counting sort in golang.

This is just for fun, maybe it could be useful to someone.

## References

- https://en.wikipedia.org/wiki/Counting_sort
- https://www.geeksforgeeks.org/counting-sort/

## Benchmarks (run them yourself, this is from a noisy laptop)

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/adamryman/counting
BenchmarkSortStdlib1-4                  10000000               105 ns/op
BenchmarkSort1-4                        20000000               130 ns/op
BenchmarkSortStdlib2-4                  10000000               260 ns/op
BenchmarkSort2-4                         5000000               219 ns/op
BenchmarkSortStdlib3-4                   5000000               261 ns/op
BenchmarkSort3-4                         5000000               246 ns/op
BenchmarkSortStdlib10-4                  2000000               796 ns/op
BenchmarkSort10-4                        2000000               583 ns/op
BenchmarkSortStdlib20-4                  1000000              1810 ns/op
BenchmarkSort20-4                        1000000              1129 ns/op
BenchmarkSortStdlib40-4                   300000              3863 ns/op
BenchmarkSort40-4                        1000000              2010 ns/op
BenchmarkSort10And1000Stdlib-4             20000             83303 ns/op
BenchmarkSort10And1000-4                   30000             41653 ns/op
BenchmarkSort10And10000Stdlib-4             2000            865794 ns/op
BenchmarkSort10And10000-4                   5000            409069 ns/op
BenchmarkSort10And100000Stdlib-4             200           8699055 ns/op
BenchmarkSort10And100000-4                   500           4333820 ns/op
BenchmarkSort10And1000000Stdlib-4             20          88007947 ns/op
BenchmarkSort10And1000000-4                   30          40521176 ns/op
BenchmarkSort100And1000000Stdlib-4            10         110538889 ns/op
BenchmarkSort100And1000000-4                  30          34794485 ns/op
BenchmarkSort1000And1000000Stdlib-4           10         141395959 ns/op
BenchmarkSort1000And1000000-4                 50          42218434 ns/op
PASS
ok      github.com/adamryman/counting   45.110s
```

## TODO

- [X] Implement counting sort
- [ ] Add useful comments
- [X] Add a test
- [ ] Add multiple tests
- [X] Add benchmarks
