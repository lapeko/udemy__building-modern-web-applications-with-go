package render

import (
	"bytes"
	"chapter3/cmd/models"
	"chapter3/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var appConfig *config.Config

func NewTemplates(config *config.Config) {
	appConfig = config
}

func ApplyDefaultRenderData(data *models.RenderData) {
	data.StringMap = map[string]string{"test": "Hello world"}
}

func Template(w http.ResponseWriter, templateName string, data *models.RenderData) {
	var templateMap map[string]*template.Template
	if appConfig.CacheTemplates {
		templateMap = appConfig.TemplateCache
	} else {
		tm, err := CreateCache()
		if err != nil {
			log.Fatal(err)
		}
		templateMap = tm
	}

	templates, exists := templateMap[templateName]

	if !exists {
		log.Fatalln("Template", templateName, " not found")
	}

	ApplyDefaultRenderData(data)

	buf := new(bytes.Buffer)
	err := templates.Execute(buf, data)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("templates/*.page.gohtml")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		temp, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("templates/*.layout.gohtml")

		if len(layouts) > 0 {
			_, err := temp.ParseGlob("templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = temp
	}

	return cache, nil
}
