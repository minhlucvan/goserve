package controllers

import (
	"net/http"
	"../models"
)

func AddHandler(res http.ResponseWriter, req *http.Request) {
	l := models.ListPage()
	t, _ := Tmpl["add"]
	t.ExecuteTemplate(res, "layout", l)
}


