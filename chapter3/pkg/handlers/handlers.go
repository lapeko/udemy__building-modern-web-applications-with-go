package handlers

import (
	"chapter3/cmd/models"
	"chapter3/pkg/config"
	"chapter3/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.Config
}

var Repo *Repository

func NewHandlers(config *config.Config) {
	Repo = &Repository{
		App: config,
	}
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.gohtml", &models.RenderData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.gohtml", &models.RenderData{})
}
