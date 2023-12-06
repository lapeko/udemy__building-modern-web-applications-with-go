package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	fmt.Printf("Server is running on port: %s", port)
	_ = http.ListenAndServe(port, nil)
}
