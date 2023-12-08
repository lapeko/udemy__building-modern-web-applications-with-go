package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func Template(w http.ResponseWriter, templateName string) {
	templates, _ := template.ParseFiles("./templates/"+templateName+".page.gohtml", "./templates/base.layout.gohtml")
	err := templates.Execute(w, nil)
	if err != nil {
		fmt.Printf("Template %s successfully send", templateName)
	}
}
