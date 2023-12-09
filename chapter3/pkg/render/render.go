package render

import (
	"fmt"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func Template(w http.ResponseWriter, templateName string) {
	templates, _ := makeTemplates(templateName)
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Printf("Template %s successfully send", templateName)
	}
}

func makeTemplates(templateName string) (*template.Template, error) {
	templates, keyExist := templateCache[templateName]
	if keyExist {
		return templates, nil
	}
	templates, err := template.ParseFiles("./templates/"+templateName+".page.gohtml", "./templates/base.layout.gohtml")
	if err != nil {
		return nil, err
	}
	templateCache[templateName] = templates
	return templates, nil
}
