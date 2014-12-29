package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		fmt.Println(v)
	}
	c <- sum
	fmt.Println("After send: ", sum)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	time.Sleep(time.Millisecond * 100)
	fmt.Println(x, y, x+y)
}
