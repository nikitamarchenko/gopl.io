/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that
reads numbers from its command-line arguments or from the standard input if there
are no arguments, and converts each number into units like temperature in Celsius
and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the
like.
*/

package main

import (
	"flag"
	"fmt"
	"gopl.io/ch2/ex2.2/conv/unitconv"
)

var temperature float64
var length float64
var weight float64

func init() {
	flag.Float64Var(&temperature, "t", 0, "Temperature")
	flag.Float64Var(&length, "l", 0, "Lenght")
	flag.Float64Var(&weight, "w", 0, "Weight")
}

func main() {
	flag.Parse()
	var printDefaults = true
	flag.Visit(func(f *flag.Flag) {
		printDefaults = false
		switch f.Name {
		case "t":
			f := unitconv.Fahrenheit(temperature)
			c := unitconv.Celsius(temperature)
			fmt.Printf("%s = %s, %s = %s\n",
				f, unitconv.FToC(f), c, unitconv.CToF(c))

		case "l":
			m := unitconv.Meters(length)
			f := unitconv.Feet(length)
			fmt.Printf("%s = %s, %s = %s\n",
				m, unitconv.MToF(m), f, unitconv.FToM(f))

		case "w":
			k := unitconv.Kilograms(weight)
			p := unitconv.Pounds(weight)
			fmt.Printf("%s = %s, %s = %s\n",
				k, unitconv.KToP(k), p, unitconv.PToK(p))
		}
	})
	if printDefaults {
		flag.PrintDefaults()
	}
}
