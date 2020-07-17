package main

import (
	"math"
	"time"
)

const (
	runes = ".,-~:;=!*#$@"

	width    = 80
	height   = 22
	gridSize = width * height

	thetaSpacing = 0.07
	phiSpacing   = 0.02

	R1 = 1
	R2 = 2
	K2 = 5
	// K1 =
)

func main() {
	var (
		A, B float64
		grid [1760]rune
		k    int
	)

	print("\x1b[2J")

	for {
		grid = setRunes(grid, ' ', gridSize)
		for theta := float64(0); theta < 2*math.Pi; theta += thetaSpacing {
			for phi := float64(0); phi < 2*math.Pi; phi += phiSpacing {
				sinPhi := math.Sin(phi)
				cosTheta := math.Cos(theta)
				sinA := math.Sin(A)
				sinTheta := math.Sin(theta)
				cosA := math.Cos(A)

				circleX := R1*cosTheta + R2
				circleY := R1 * sinTheta

				D := R1 / (sinPhi*circleX*sinA + sinTheta*cosA + K2)

				l := math.Cos(phi)
				cosB := math.Cos(B)
				sinB := math.Sin(B)
				t := sinPhi*circleX*cosA - circleY*sinA

				x := int(40 + 30*D*(l*circleX*cosB-t*sinB))
				y := int(12 + 15*D*(l*circleX*sinB+t*cosB))
				index := x + width*y
				N := int(8 * ((sinTheta*sinA-sinPhi*cosTheta*cosA)*cosB - sinPhi*cosTheta*sinA - sinTheta*cosA - l*cosTheta*sinB))

				if height > y && y > 0 && x > 0 && width > x {
					if N > 0 {
						grid[index] = rune(runes[N])
						continue
					}

					grid[index] = '.'
				}
			}
		}

		print("\x1b[H")
		for k = 0; k < gridSize; k++ {
			print(string(grid[k]))
			A = A + 0.00004
			B = B + 0.00001
		}
		time.Sleep(30 * time.Millisecond)
	}
}

func setRunes(dst [gridSize]rune, value rune, n int) [gridSize]rune {
	for i := 0; i < n; i++ {
		if i%width == width-1 {
			dst[i] = '\n'
			continue
		}

		dst[i] = value
	}

	return dst
}
