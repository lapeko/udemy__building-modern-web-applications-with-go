package handlers

import (
	"chapter3/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about")
}
