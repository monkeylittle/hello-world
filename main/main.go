package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Fabio!")
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", internalServerError)
	http.ListenAndServe(":8000", nil)
}
