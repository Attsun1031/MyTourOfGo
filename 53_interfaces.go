package main

import (
	"fmt"
	"math"
	"reflect"
)

type Abser interface {
	Abs() (z float64)
}

type MyFloat float64

func (f MyFloat) Abs() (value float64) {
	if f < 0 {
		value = float64(-f)
	}
	value = float64(f)
	return
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// a = v
	fmt.Println(a.Abs())
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(v))
}
