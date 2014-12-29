package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	values := []int{0, 2, 4, 5}
	for i, v := range values {
		fmt.Println(i, v)
	}
}
