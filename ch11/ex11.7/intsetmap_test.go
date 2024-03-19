package intset

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

type IntSetMap struct {
	m map[int]bool
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSetMap) Has(x int) bool {
	return s.m[x]
}

// Add adds the non-negative value x to the set.
func (s *IntSetMap) Add(x int) {
	if s.m == nil {
		s.m = make(map[int]bool)
	}
	s.m[x] = true
}

// UnionWith sets s to the union of s and t.
func (s *IntSetMap) UnionWith(t IntSetMap) {
	for v := range t.m {
		s.m[v] = true
	}
}

func (s *IntSetMap) String() string {
	arr := make([]int, 0, len(s.m))
	arr2 := make([]string, len(s.m))
	for v := range s.m {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	for i, v := range arr {
		arr2[i] = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(arr2, " "))
}

func benchmarkIntSetMapHas() func() {
	r := rand.New(rand.NewSource(seed))
	var x IntSetMap
	x.Add(r.Intn(maxInt))
	h := r.Intn(maxInt)
	x.Add(h)
	return func() {
		x.Has(h)
	}
}

func benchmarkIntSetMapAdd() func() {
	r := rand.New(rand.NewSource(seed))
	h := r.Intn(maxInt)
	return func() {
		var x IntSetMap
		x.Add(h)
	}
}

func benchmarkIntSetMapUnion() func() {
	r := rand.New(rand.NewSource(seed))
	h1 := r.Intn(maxInt)
	h2 := r.Intn(maxInt)
	return func() {
		var x, y IntSetMap
		x.Add(h1)
		y.Add(h2)
		x.UnionWith(y)
	}
}

func BenchmarkIntSetMap_Has1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapHas, 1)
}

func BenchmarkIntSetMap_Has10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapHas, 10)
}

func BenchmarkIntSetMap_Has100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapHas, 100)
}

func BenchmarkIntSetMap_Has1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapHas, 1000)
}

func BenchmarkIntSetMap_Add1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapAdd, 1)
}

func BenchmarkIntSetMap_Add10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapAdd, 10)
}

func BenchmarkIntSetMap_Add100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapAdd, 100)
}

func BenchmarkIntSetMap_Add1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapAdd, 1000)
}

func BenchmarkIntSetMap_Union1(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapUnion, 1)
}

func BenchmarkIntSetMap_Union10(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapUnion, 10)
}

func BenchmarkIntSetMap_Union100(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapUnion, 100)
}

func BenchmarkIntSetMap_Union1000(b *testing.B) {
	benchmarkN(b, benchmarkIntSetMapUnion, 1000)
}
