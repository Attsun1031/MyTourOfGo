package main

import (
	"fmt"
	"math"
)

var i int = 1

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (a, b string) {
	a = y + "hoge"
	b = x
	return
}

func newton(x, z float64) float64 {
	return z - ((math.Pow(z, 2) - x) / (z * 2))
}

func NewtonSqrt(x float64) (z float64) {
	z = 1.0
	previous := 0.0
	for i := 0; i < 10; i++ {
		z = newton(x, z)

		if z == previous {
			break
		} else {
			previous = z
		}
	}
	return z
}

func main() {
	for x := 0; x < 10; x++ {
		fx := float64(x)
		fmt.Printf("Input: %d, math.Sqrt: %f, NewtonSqrt: %f\n", x, math.Sqrt(fx), NewtonSqrt(fx))
	}
}
