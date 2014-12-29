package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case now, _ := <-tick:
			fmt.Println("tick ", now)
		case now, _ := <-boom:
			fmt.Println("boom", now)
		}
	}
}
