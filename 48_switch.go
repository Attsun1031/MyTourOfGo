package main

import (
	"fmt"
	"math/cmplx"
)

func newton(x, z complex128) complex128 {
	return z - ((cmplx.Pow(z, 3) - x) / (3*cmplx.Pow(z, 2)))
}

func Cbrt(x complex128) (z complex128) {
	z = 1
	var previous complex128 = 0
	for i := 0; i < 10; i++ {
		z = newton(x, z)

		if z == previous {
			break
		} else {
			previous = z
		}
	}
	return
}

func main() {
	fmt.Printf("Cbrt: %v\n", Cbrt(2))
	fmt.Printf("cmplx.Pow: %v\n", cmplx.Pow(2, 3))
}
