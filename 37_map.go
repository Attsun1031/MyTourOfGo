package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	key := "Bell Labs"
	key2 := "松が谷"
	m[key] = Vertex{
		40.6833, -74.3977,
	}
	m[key2] = Vertex{
		Long:30.2002, Lat:-10.1010,
	}
	fmt.Println(m[key])
	fmt.Println(m[key2])
}
