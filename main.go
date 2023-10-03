package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	err:=r.ParseForm() 
	if err!=nil{
		fmt.Fprintf(w, "parseform error ") 
		return 
	}
	fmt.Fprintf(w, "post request successfull!! ")
	name:=r.FormValue("name") 
	address:=r.FormValue("address") 
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found ", http.StatusNotFound)
		return 
	} 

	if r.Method!="GET"{
		http.Error(w,"method is not supported", http.StatusNotFound)
		return 
	}

	fmt.Fprintf(w,"Hello!") 
}

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")

	err:=http.ListenAndServe(":8080", nil)
	if err!=nil{
		log.Fatal(err )
	}
}