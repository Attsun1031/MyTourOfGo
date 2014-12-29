package main

import "fmt"

// 関数を値にすると再帰呼び出しができない。
func getFibonacci (x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return getFibonacci(x - 1) + getFibonacci(x - 2)
	}
}

func fibonacci() func() int {
	count := 0
	return func() int {
		retval := getFibonacci(count)
		count += 1
		return retval
	}
}

func main() {
	f := fibonacci()
	for i := 0;  i < 10; i++ {
		fmt.Println(f())
	}
}
