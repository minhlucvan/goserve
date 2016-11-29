package controllers

import (
	"net/http"
	"../models"
)

func EditHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		p = models.StringPage{Title: title}
	}
	t, _ := Tmpl["edit"]
	t.ExecuteTemplate(res, "layout", p)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(res, req, "/", http.StatusMovedPermanently)
		return
	}
	req.ParseForm()
	title := req.FormValue("title")

	if len(title) == 0{
		title = req.URL.Path[len("/save/"):]
	}

	body := req.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	//page, _ := getPageFromReq(req)
	p.Save()
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}