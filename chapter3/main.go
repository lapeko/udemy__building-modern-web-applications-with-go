package main

import (
	"errors"
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

func divideHandler(w http.ResponseWriter, r *http.Request) {
	a := 10.0
	b := 2.0
	result, err := divideNumber(a, b)
	if err != nil {
		_, _ = fmt.Fprintf(w, fmt.Sprintf(err.Error()))
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f / %f is %f", a, b, result))
}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}

func divideNumber(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0 is forbidden to use as a divider")
	}
	return a / b, nil
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/divide", divideHandler)
	fmt.Printf("Server is running on port: %s", port)
	_ = http.ListenAndServe(port, nil)
}
