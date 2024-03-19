/*

Exercise 11.7: Write benchmarks for Add, UnionWith, and other methods of
*IntSet (§6.5) using large pseudo-random inputs. How fast can you make these
methods run? How does the choice of word size affect performance? How fast is
IntSet compared to a set implementation based on the built-in map type?

32bit version is twice slower.
Has faster on IntSet, Add and Union on IntSetMap.

go test -bench=.
goos: linux
goarch: amd64
pkg: gopl.io/ch11/ex11.7
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkIntSet_Has1-8           	465163972	        2.521 ns/op
BenchmarkIntSet_Has10-8          	58542085	       20.60 ns/op
BenchmarkIntSet_Has100-8         	5881375	      209.6 ns/op
BenchmarkIntSet_Has1000-8        	 595796	     2108 ns/op
BenchmarkIntSet_Add1-8           	5029764	      246.4 ns/op
BenchmarkIntSet_Add10-8          	 455716	     2539 ns/op
BenchmarkIntSet_Add100-8         	  46238	    23187 ns/op
BenchmarkIntSet_Add1000-8        	   4839	   234840 ns/op
BenchmarkIntSet_Union1-8         	1691985	      797.8 ns/op
BenchmarkIntSet_Union10-8        	 146869	     8184 ns/op
BenchmarkIntSet_Union100-8       	  14694	    81672 ns/op
BenchmarkIntSet_Union1000-8      	   1472	   725846 ns/op
BenchmarkIntSetMap_Has1-8        	223774202	        5.354 ns/op
BenchmarkIntSetMap_Has10-8       	25719264	       49.33 ns/op
BenchmarkIntSetMap_Has100-8      	2212196	      490.9 ns/op
BenchmarkIntSetMap_Has1000-8     	 248143	     4684 ns/op
BenchmarkIntSetMap_Add1-8        	16677982	       72.58 ns/op
BenchmarkIntSetMap_Add10-8       	1637276	      737.5 ns/op
BenchmarkIntSetMap_Add100-8      	 144379	     7806 ns/op
BenchmarkIntSetMap_Add1000-8     	  16972	    70761 ns/op
BenchmarkIntSetMap_Union1-8      	4503232	      280.0 ns/op
BenchmarkIntSetMap_Union10-8     	 382154	     2855 ns/op
BenchmarkIntSetMap_Union100-8    	  45912	    26843 ns/op
BenchmarkIntSetMap_Union1000-8   	   3898	   308339 ns/op

GOARCH=386 go test -bench=.
goos: linux
goarch: 386
pkg: gopl.io/ch11/ex11.7
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkIntSet_Has1-8           	335706619	        3.484 ns/op
BenchmarkIntSet_Has10-8          	46177867	       25.68 ns/op
BenchmarkIntSet_Has100-8         	4691703	      255.3 ns/op
BenchmarkIntSet_Has1000-8        	 480876	     2477 ns/op
BenchmarkIntSet_Add1-8           	3068821	      365.6 ns/op
BenchmarkIntSet_Add10-8          	 328502	     3715 ns/op
BenchmarkIntSet_Add100-8         	  29044	    38647 ns/op
BenchmarkIntSet_Add1000-8        	   3202	   367689 ns/op
BenchmarkIntSet_Union1-8         	 846296	     1374 ns/op
BenchmarkIntSet_Union10-8        	  89437	    13525 ns/op
BenchmarkIntSet_Union100-8       	   8218	   133858 ns/op
BenchmarkIntSet_Union1000-8      	    793	  1328100 ns/op
BenchmarkIntSetMap_Has1-8        	172868709	        6.200 ns/op
BenchmarkIntSetMap_Has10-8       	22957153	       52.96 ns/op
BenchmarkIntSetMap_Has100-8      	2047446	      514.4 ns/op
BenchmarkIntSetMap_Has1000-8     	 231925	     4942 ns/op
BenchmarkIntSetMap_Add1-8        	10798816	      110.1 ns/op
BenchmarkIntSetMap_Add10-8       	 908619	     1102 ns/op
BenchmarkIntSetMap_Add100-8      	 111208	    11040 ns/op
BenchmarkIntSetMap_Add1000-8     	  10000	   106335 ns/op
BenchmarkIntSetMap_Union1-8      	3057190	      433.8 ns/op
BenchmarkIntSetMap_Union10-8     	 268660	     3944 ns/op
BenchmarkIntSetMap_Union100-8    	  27814	    40708 ns/op
BenchmarkIntSetMap_Union1000-8   	   2673	   395899 ns/op
PASS
ok  	gopl.io/ch11/ex11.7	31.942s

*/

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"testing"

	"math/rand"
)

var seed int64 = 97580237

var maxInt = 100 * 32

func benchmarkIntSetHas() func() {
	r := rand.New(rand.NewSource(seed))
	var x IntSet
	x.Add(r.Intn(maxInt))
	h := r.Intn(maxInt)
	x.Add(h)
	return func() {
		x.Has(h)
	}
}

func benchmarkIntSetAdd() func() {
	r := rand.New(rand.NewSource(seed))
	h := r.Intn(maxInt)
	return func() {
		var x IntSet
		x.Add(h)
	}
}

func benchmarkIntSetUnion() func() {
	r := rand.New(rand.NewSource(seed))
	h1 := r.Intn(maxInt)
	h2 := r.Intn(maxInt)
	return func() {
		var x, y IntSet
		x.Add(h1)
		y.Add(h2)
		x.UnionWith(&y)
	}
}

func benchmarkN(b *testing.B, f func() func(), n int) {
	run := f()
	for i := 0; i < b.N; i++ {
		for ii := 0; ii < n; ii++ {
			run()
		}
	}
}

func BenchmarkIntSet_Has1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetHas, 1)
}

func BenchmarkIntSet_Has10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetHas, 10)
}

func BenchmarkIntSet_Has100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetHas, 100)
}

func BenchmarkIntSet_Has1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetHas, 1000)
}

func BenchmarkIntSet_Add1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetAdd, 1)
}

func BenchmarkIntSet_Add10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetAdd, 10)
}

func BenchmarkIntSet_Add100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetAdd, 100)
}

func BenchmarkIntSet_Add1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetAdd, 1000)
}

func BenchmarkIntSet_Union1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetUnion, 1)
}

func BenchmarkIntSet_Union10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetUnion, 10)
}

func BenchmarkIntSet_Union100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetUnion, 100)
}

func BenchmarkIntSet_Union1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetUnion, 1000)
}
