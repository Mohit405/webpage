package main

import (
	"net/http"
	// "text/template"
)
func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
