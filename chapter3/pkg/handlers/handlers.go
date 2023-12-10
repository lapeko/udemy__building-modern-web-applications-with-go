package handlers

import (
	"chapter3/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.gohtml")
}
