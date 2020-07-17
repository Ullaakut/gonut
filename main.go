package main

import (
	"math"
	"time"
)

const runes = ".,-~:;=!*#$@"

func main() {
	var (
		A, B, i, j float64
		// z          [1760]float64
		b [1760]rune
		k int
	)

	print("\x1b[2J")

	for {
		b = memsetRunes(b, ' ', 1760)
		for j = 0; j < 6.28; j += 0.07 {
			for i = 0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)

				h := d + 2
				D := 1 / (c*h*e + f*g + 5)

				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e

				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := x + 80*y
				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))

				if 22 > y && y > 0 && x > 0 && 80 > x { // && D > z[o] {
					// z[o] = D

					if N > 0 {
						b[o] = rune(runes[N])
					} else {
						b[o] = '.'
					}
				}
			}
		}

		print("\x1b[H")
		for k = 0; k < 1760; k++ {
			print(string(b[k]))
			A = A + 0.00004
			B = B + 0.00001
		}
		time.Sleep(30 * time.Millisecond)
	}
}

func memsetRunes(dst [1760]rune, value rune, n int) [1760]rune {
	for i := 0; i < n; i++ {
		if i%80 == 79 {
			dst[i] = '\n'
			continue
		}

		dst[i] = value
	}

	return dst
}
