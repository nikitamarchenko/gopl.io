/*
Exercise 3.5: Implement a full-color Mandelbrot set using the function image.NewRGBA
and the type color.RGBA or color.YCbCr
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

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// img := image.NewRGBA(image.Rect(0, 0, width, height))

	img := image.NewCMYK(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	f, err := os.OpenFile("out.png", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return
	}

	png.Encode(f, img)

	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 360 * 10
	const contrast = 4

	var v complex128
	var n uint64 = 0
	for ; n < iterations && cmplx.Abs(v) <= 2; n++ {
		v = v*v + z
	}

	//from https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set
	//hsl = [powf((i / max) * 360, 1.5) % 360, 50, (i / max) * 100]

	if n == iterations {
		return color.Black
	}

	return toRGB(
		math.Mod(
			math.Pow(
				(float64(n)/float64(iterations))*360.0,
				1.5),
			360.0),
		50,
		((float64(n)/float64(iterations))*100),
	)
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
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

	// fmt.Println(h)

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
