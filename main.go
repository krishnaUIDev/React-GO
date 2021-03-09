package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage request")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fprintf(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequest()
}