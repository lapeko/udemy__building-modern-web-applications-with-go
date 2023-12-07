package main

import (
	"chapter3/pkg/handlers"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	fmt.Printf("Server is running on port: %s", port)
	_ = http.ListenAndServe(port, nil)
}
