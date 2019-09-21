package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Printf("First module")
	http.HandleFunc("/", myFunc)
}

func myFunc(w http.ResponseWriter, r *http.Request) {
	strings.Compare("a", "b")
}
