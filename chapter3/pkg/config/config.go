package config

import "text/template"

type Config struct {
	CacheTemplates bool
	TemplateCache  map[string]*template.Template
}
