package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	} else {
		return f, nil
	}
}

func main() {
	//fmt.Println(Sqrt(2))
	_, err := Sqrt(-2)
	fmt.Println(err)
}
