package main

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about")
}
