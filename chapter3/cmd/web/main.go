package main

import (
	"chapter3/pkg/config"
	"chapter3/pkg/handlers"
	"chapter3/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	var appConfig config.Config
	templateCache, err := render.CreateCache()

	if err != nil {
		log.Fatal(templateCache)
	}

	appConfig = config.Config{
		TemplateCache:  templateCache,
		CacheTemplates: true,
	}

	render.NewTemplates(&appConfig)
	handlers.NewHandlers(&appConfig)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Server is running on port: %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
