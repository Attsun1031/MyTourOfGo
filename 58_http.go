package main

import (
	"fmt"
	"net/http"
	"strings"
)

type String string

type Struct struct {
	Greeting string
	Pucnt    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(s))
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := []string{r.Form["x"][0], s.Greeting, s.Pucnt, s.Who}
	fmt.Fprint(w, strings.Join(values, ","))
}

func main() {
	str := String("xxx")
	stc := &Struct{"greet", "puct", "who"}
	http.Handle("/string", str)
	http.Handle("/struct", stc)
	http.ListenAndServe("localhost:4000", nil)
}
