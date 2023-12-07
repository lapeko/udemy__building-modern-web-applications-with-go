package handlers

import (
	"chapter3/pkg/render"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
