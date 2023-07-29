package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("form", r.Form)
	fmt.Fprintf(w, "Hello astaxie!")
}

func createServer() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Create server error: ", err)
	} else {
		fmt.Println("Server is running on port 9090")
	}
}
