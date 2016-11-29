package controllers

import (
	"html/template"
	"../../config"
)

var Tmpl = map[string]*template.Template{
	"home": template.Must(template.ParseFiles(config.VIEW_PATH + "layout.html", config.VIEW_PATH + "index.html")),
	"edit": template.Must(template.ParseFiles(config.VIEW_PATH + "layout.html", config.VIEW_PATH + "edit.html")),
	"add": template.Must(template.ParseFiles(config.VIEW_PATH + "layout.html", config.VIEW_PATH + "add.html")),
}