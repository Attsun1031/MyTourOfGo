package main

import (
	"fmt"
	"os"
	"compress/gzip"
	"io/ioutil"
)

func main() {
	f, _ := os.Open("test.txt.gz")
	gzf, _ := gzip.NewReader(f)
	contents, _ := ioutil.ReadAll(gzf)
	fmt.Printf(string(contents))
}
