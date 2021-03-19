package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, greet string) {
	fmt.Fprint(w, greet)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Gra Gra\n")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
