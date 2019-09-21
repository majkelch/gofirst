package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("First module")
	http.HandleFunc("/", myFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Response struct {
	FirstString string `json:"first"`
	SecondString string `json:"second"`
	CompareResult int `json:"result"`
	Executed time.Time `json:"executed"`
}

func myFunc(w http.ResponseWriter, r *http.Request) {
	const first string = "some string"
	const second string = "another string"
	comp := compareStrings(first, second)
	res := Response{
		first,
		second,
		comp,
		time.Now(),
	}
	js, err := json.Marshal(res)
	fmt.Println(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func compareStrings(a string, b string) int {
	return strings.Compare(a, b)
}
