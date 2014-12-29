package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (ret map[string]int) {
	ret = make(map[string]int)
	for _, chr := range strings.Fields(s) {
	//for _, chr := range strings.Split(s, " ") {
		ret[string(chr)] += 1
	}
	return
}

func main() {
	wc.Test(WordCount)
}
