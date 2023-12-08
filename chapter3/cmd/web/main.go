package main

import (
	"chapter3/pkg/handlers"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Server is running on port: %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
