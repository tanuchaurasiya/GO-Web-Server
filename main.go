package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style = 'color:steelblue'> hello world</h1>"))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := []Book{
		{
			Title:  "this is title1",
			Author: "this is author1",
			Pages:  100},
		{
			Title:  "this is title2",
			Author: "this is author2",
			Pages:  200,
		},
	}

	json.NewEncoder(w).Encode(book)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from main"))
	})
	http.HandleFunc("/getbook", getBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
