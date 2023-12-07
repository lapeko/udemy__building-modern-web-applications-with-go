package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	templates, _ := template.ParseFiles("./templates/" + templateName + ".html")
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Printf("Template %s successfully send", templateName)
	}
}
