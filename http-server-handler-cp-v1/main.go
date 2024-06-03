package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		currentTime := time.Now()
		response := fmt.Sprintf("%v, %v %v %v", currentTime.Weekday(), currentTime.Day(), currentTime.Month(), currentTime.Year())
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte(response))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
