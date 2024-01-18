/*
Exercise 5.16: Write a variadic version of strings.Join.
*/

package main

import "strings"

func strJoin(sep string, values ...string) string {
	
	switch len(values) {
	case 0:
		return ""
	case 1:
		return values[0]
	}

	var b strings.Builder

	b.WriteString(values[0])
	for _, s := range values[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	return b.String()
}
