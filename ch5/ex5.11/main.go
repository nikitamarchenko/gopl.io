/*
Exercise 5.11: The instructor of the linear algebra course decides that
calculus is now a prerequisite. Extend the topoSort function to report cycles.
*/

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {
		"data structures": true},
	"calculus": {
		"linear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},
	"data structures": {
		"discrete math": true},
	"databases": {
		"data structures": true},
	"discrete math": {
		"intro to programming": true,
	},
	"formal languages": {
		"discrete math": true},
	"networks": {
		"operating systems": true},
	"operating systems": {
		"data structures":       true,
		"computer organization": true,
	},
	"programming languages": {
		"data structures":       true,
		"computer organization": true},
	"linear algebra": {
		"calculus": true},
}

//!-table

// !+main
func main() {
	courses, cycles := topoSort(prereqs)
	fmt.Println("Courses:")
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	fmt.Println("Cycles:")
	for _, cycle := range cycles {

		for i, c := range cycle {
			fmt.Printf("%*s- %s\n", i*2, "", c)
		}
	}
}

func topoSort(m map[string]map[string]bool) ([]string, [][]string) {
	var order []string
	seen := make(map[string]bool)
	trace := make(map[string]bool)
	path := []string{}
	var sortVisitAll func(items map[string]bool)
	cycles := [][]string{}
	depth := -1
	sortVisitAll = func(items map[string]bool) {
		depth++
		for item := range items {
			//fmt.Printf("%*s- %s\n", depth*2, "", item)
			if _, ok := trace[item]; ok {
				//fmt.Println("***found cycle", path, item)
				p := make([]string, len(path)+1)
				copy(p, path)
				p[len(path)] = item
				cycles = append(cycles, p)

			} else {
				trace[item] = true
				path = append(path, item)
				if !seen[item] {
					seen[item] = true
					sortVisitAll(m[item])
					order = append(order, item)
				}
				delete(trace, item)
				path = path[:len(path)-1]
			}
		}
		depth--
	}

	var keys = make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	sortVisitAll(keys)

	for i, c := range cycles {
		for len(c) > 1 && c[0] != c[len(c)-1] {
			c = c[1:]
		}
		cycles[i] = c
	}

	return order, cycles
}

//!-main
