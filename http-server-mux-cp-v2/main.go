package main

import (
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		formattedTime := currentTime.Format("Monday, 02 January 2006")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(formattedTime))
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello there"))
		} else {
			message := "Hello, " + name + "!"
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(message))
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
