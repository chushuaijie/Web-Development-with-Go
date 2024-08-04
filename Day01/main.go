package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome sites!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Start the server on 300...")
	http.ListenAndServe(":3000", nil)
}
