package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tanu"))
	})
	http.ListenAndServe(":8080", nil)
}
