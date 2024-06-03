package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Fprintf(w, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Hello there")
		} else {
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	}) // TODO: replace this
}

func main() {
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	// TODO: answer here
	http.ListenAndServe("localhost:8080", nil)
}
