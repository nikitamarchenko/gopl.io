/*
Exercise 7.6: Add support for Kelvin temperatures to tempflag.
*/

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch7/ex7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
