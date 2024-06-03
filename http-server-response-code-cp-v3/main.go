package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		exist := IsNameExists(name)
		if !exist {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Name is exists"))
	}
	// TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
