package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Bell Labs": {
			40.6833, -74.39967,
		},
		"Google": {
			37.4203, -122.08408,
		},
	}
	fmt.Println(m)
}
