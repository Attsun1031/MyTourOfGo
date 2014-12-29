package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello\n")
}
