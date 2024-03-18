/*

Exercise 11.6: Write benchmarks to compare the PopCount implementation in
Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5. At what
point does the table-based approach break even?

2.6.2 -> gopl.io/ch2/popcount

ex in gopl.io/ch2/ex2.3+4+5

result:

go test -bench=.
goos: linux
goarch: amd64
pkg: gopl.io/ch11/ex11.6
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkPopCount10-8               	30504496	   32.89 ns/op
BenchmarkPopCount100-8              	3515758	      328.8 ns/op
BenchmarkPopCount1000-8             	 372631	     3236 ns/op
BenchmarkPopCount2_10-8             	10165023	  119.4 ns/op
BenchmarkPopCount2_100-8            	1000000	     1149 ns/op
BenchmarkPopCount2_1000-8           	 106720	    12287 ns/op
BenchmarkPopCountByClearing10-8     	9308241	      129.2 ns/op
BenchmarkPopCountByClearing100-8    	 934052	     1230 ns/op
BenchmarkPopCountByClearing1000-8   	  98002	    12459 ns/op
BenchmarkPopCountByShifting10-8     	2610063	      442.1 ns/op
BenchmarkPopCountByShifting100-8    	 281937	     4252 ns/op
BenchmarkPopCountByShifting1000-8   	  28118	    43116 ns/op
PASS
ok  	gopl.io/ch11/ex11.6	16.061s

At what point does the table-based approach break even?
a Over 1k.

*/

package ex116_test

import (
	"testing"

	"gopl.io/ch2/ex2.3+4+5/popcount"
)

func benchmark(b *testing.B, f func(x uint64) int, size int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			f(uint64(0x1234567890ABCDEF))
		}
	}
}

var pc [256]byte

func benchmarkTable(b *testing.B, f func(x uint64) int, size int) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			f(uint64(0x1234567890ABCDEF))
		}
	}
}

func BenchmarkPopCount10(b *testing.B) {
	benchmarkTable(b, popcount.PopCount, 10)
}

func BenchmarkPopCount100(b *testing.B) {
	benchmarkTable(b, popcount.PopCount, 100)
}

func BenchmarkPopCount1000(b *testing.B) {
	benchmarkTable(b, popcount.PopCount, 1000)
}

func BenchmarkPopCount2_10(b *testing.B) {
	benchmarkTable(b, popcount.PopCount2, 10)
}

func BenchmarkPopCount2_100(b *testing.B) {
	benchmarkTable(b, popcount.PopCount2, 100)
}

func BenchmarkPopCount2_1000(b *testing.B) {
	benchmarkTable(b, popcount.PopCount2, 1000)
}

func BenchmarkPopCountByClearing10(b *testing.B) {
	benchmark(b, popcount.PopCountByClearing, 10)
}

func BenchmarkPopCountByClearing100(b *testing.B) {
	benchmark(b, popcount.PopCountByClearing, 100)
}

func BenchmarkPopCountByClearing1000(b *testing.B) {
	benchmark(b, popcount.PopCountByClearing, 1000)
}

func BenchmarkPopCountByShifting10(b *testing.B) {
	benchmark(b, popcount.PopCountByShifting, 10)
}

func BenchmarkPopCountByShifting100(b *testing.B) {
	benchmark(b, popcount.PopCountByShifting, 100)
}

func BenchmarkPopCountByShifting1000(b *testing.B) {
	benchmark(b, popcount.PopCountByShifting, 1000)
}
