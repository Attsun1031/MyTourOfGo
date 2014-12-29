package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	elm := m["Answer"]
	fmt.Println(elm)

	delete(m, "Answer")
	elm, ok := m["Answer"]
	fmt.Println(elm, ok)
}
