package main

import (
	"fmt"
	"net/http"
)

func main() {
	// The Server must serve on some endpoint. Endpoint states the resource that has to be accessed
	// Endpoint could be homepage of the server
	// Use HandleFunc will register the handler function for the given pattern
	// handler function handles the requests coming at a given endpoint
	// In nut-shell, Handler Func ties request to a specific endpoint to a particular Handler Function

	http.HandleFunc("/", homepage) // homepage is the handler function

	http.ListenAndServe("localhost:10000", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Welcome to the HomePage<h1><br><h3>Did you like it ?</h3>")

	if err != nil {
		fmt.Println("Err writing to the io writer\t", err)
		return
	}
	fmt.Println("Name of the endpoint hit : HomePage") // Write to the console
	return
}
