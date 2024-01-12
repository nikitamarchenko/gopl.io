package main

import (
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {

	res := topoSort(prereqs)

	getPrereqs := func(s string) map[string]bool {
		r, ok := prereqs[s]
		if !ok {
			return map[string]bool{}
		}

		keys := make(map[string]bool)
		for k := range r {
			keys[k] = true
		}
		return keys
	}

	for i, r := range res {
		pr := getPrereqs(r)
		for ii := i - 1; ii >= 0; ii-- {
			if len(pr) == 0 {
				break
			}
			delete(pr, res[ii])
		}
		if len(pr) > 0 {
			fmt.Printf("error %v\n", pr)
		}
	}
}
