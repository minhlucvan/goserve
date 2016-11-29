package controllers

import (
	"net/http"
	"../models"
)

func IndexHander(res http.ResponseWriter, req *http.Request) {
	l := models.ListPage()
	t, _ := Tmpl["home"]
	t.ExecuteTemplate(res, "layout", l)
}
