/*
Exercise 8.5: Take an existing CPU-bound sequential program, such as the
Mandelbrot program of Section 3.3 or the 3-D surface computation of
Section 3.2, and execute its main loop in parallel using channels for
communication. How much faster does it run on a multiprocessor machine? What
is the optimal number of goroutines to use?

*/

package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	p := flag.Int("p", 1, "parallelism")
	flag.Parse()
	fmt.Print(calc(*p))
}

func calc(p int) string {
	result := [cells * cells]string{}
	c := make(chan int, cells*cells)
	var wg sync.WaitGroup
	for i := 0; i < p; i++ {
		wg.Add(1)
		go func() {
			for a := range c {
				i := a / cells
				j := a % cells
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				result[i*100+j] = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
			wg.Done()
		}()
	}
	wg.Add(1)
	go func() {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				c <- i*cells + j
			}
		}
		close(c)
		wg.Done()
	}()

	wg.Wait()

	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height))

	for _, v := range result {
		b.WriteString(v)
	}
	b.WriteString("</svg>")

	return b.String()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
