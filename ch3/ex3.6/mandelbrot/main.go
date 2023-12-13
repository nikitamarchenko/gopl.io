/*
Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by
computing the color value at several points within each pixel and taking the average.
The simplest method is to divide each pixel into four “subpixels.” Implement it.
*/

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

const iterations = 360 * 10

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var data [width][height]uint64

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			data[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {

			var points [5]color.RGBA

			up_x, up_y := px, py+1
			if up_y >= height {
				up_y = py
			}
			points[0] = getColorFromMandelbrot(data[up_x][up_y])

			down_x, down_y := px, py-1
			if down_y < 0 {
				down_y = 0
			}
			points[1] = getColorFromMandelbrot(data[down_x][down_y])

			left_x, left_y := px-1, py
			if left_x < 0 {
				left_x = 0
			}
			points[2] = getColorFromMandelbrot(data[left_x][left_y])

			right_x, right_y := px+1, py
			if right_x >= width {
				right_x = px
			}
			points[3] = getColorFromMandelbrot(data[right_x][right_y])
			points[4] = getColorFromMandelbrot(data[px][py])

			var r, g, b uint64

			for _, c := range points {
				r += uint64(c.R)
				g += uint64(c.G)
				b += uint64(c.B)
			}

			img.Set(px, py, color.RGBA{uint8(r / 5), uint8(g / 5), uint8(b / 5), 0xff})
		}
	}

	f, err := os.OpenFile("out.png", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return
	}

	png.Encode(f, img)
}

func getColorFromMandelbrot(n uint64) color.RGBA {

	//from https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set
	//hsl = [powf((i / max) * 360, 1.5) % 360, 50, (i / max) * 100]

	if n == iterations {
		return color.RGBA{0x0, 0x0, 0x0, 0xff}
	}

	return toRGB(
		math.Mod(
			math.Pow(
				(float64(n)/float64(iterations))*360.0,
				1.5),
			360.0),
		50,
		((float64(n) / float64(iterations)) * 100),
	)
}

func mandelbrot(z complex128) uint64 {
	var v complex128
	var n uint64 = 0
	for ; n < iterations && cmplx.Abs(v) <= 2; n++ {
		v = v*v + z
	}
	return n
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}

func toRGB(h, s, l float64) color.RGBA {

	if s == 0 {
		// it's gray
		return color.RGBA{uint8(l), uint8(l), uint8(l), 0xff}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
