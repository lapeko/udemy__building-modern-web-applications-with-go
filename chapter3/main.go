package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("About page. The sum of 2 and 2 is %d", calculateSum(2, 2)))
}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Printf("Server is running on port: %s", port)
	_ = http.ListenAndServe(port, nil)
}
