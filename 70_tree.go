package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
	"sort"
)

func doWalk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		doWalk(t.Left, ch)
	}
	if t.Right != nil {
		doWalk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	a1 := make([]int, 0, 100)
	a2 := make([]int, 0, 100)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for v := range c1 {
		a1 = append(a1, v)
	}
	for v := range c2 {
		a2 = append(a2, v)
	}
	result := testEq(a1, a2)
	return result
}

func testEq(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func makeTree() *tree.Tree {
	t := &tree.Tree{
		Value: 3,
		Left: &tree.Tree{
			Value: 1,
			Left: &tree.Tree{
				Value: 1,
			},
			Right: &tree.Tree{
				Value: 2,
			},
		},
		Right: &tree.Tree{
			Value: 8,
			Left: &tree.Tree{
				Value: 5,
			},
			Right: &tree.Tree{
				Value: 13,
			},
		},
	}
	return t
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("t1: ", t1.String())
	fmt.Println("t2: ", t2.String())
	fmt.Println(Same(t1, t2))
}
