package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templateCache, err = createCache()

func Template(w http.ResponseWriter, templateName string) {
	if err != nil {
		log.Fatal(err)
	}
	cachedTemplate, exists := templateCache[templateName]

	if !exists {
		log.Fatal("Template ", templateName, " does not exist")
	}
	buf := new(bytes.Buffer)
	err = cachedTemplate.Execute(buf, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}
}

func createCache() (map[string]*template.Template, error) {
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
