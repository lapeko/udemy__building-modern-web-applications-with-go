package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	templates, _ := template.ParseFiles("./templates/" + templateName + ".html")
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Printf("Template %s successfully send", templateName)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Printf("Server is running on port: %s", port)
	_ = http.ListenAndServe(port, nil)
}
