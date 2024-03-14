package intset

import (
	"fmt"
	"sort"
	"strings"
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
