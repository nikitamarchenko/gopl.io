package main

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkCalc(b *testing.B) {
	var table = []struct{ input int }{}

	cpu := runtime.NumCPU()
	for i := 1; i < cpu; i++ {
		table = append(table, struct{ input int }{i})
	}
	for i := 1; i <= 8; i++ {
		table = append(table, struct{ input int }{i * cpu})
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("calc(%d)", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				calc(v.input)
			}
		})
	}
}
