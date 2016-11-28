package main

import (
	"net/http"
	"io/ioutil"
	"html/template"
	"fmt"
)

type Page struct {
	Title string
	Body []byte
}

type StringPage struct {
	Title string
	Body string
}

func parsePage(p *Page) StringPage {
	return StringPage{Title: p.Title, Body: string(p.Body)}
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func handler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprint(res, "404 not found")
		return
	}
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	p, _ := loadPage("data")
	t, _ := template.ParseFiles("view/layout.html", "view/index.html")
	t.ExecuteTemplate(res, "layout", parsePage(p))
}

func editHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("view/edit.html")
	t.Execute(res, parsePage(p))
}

func saveHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	body := req.Form["body"][0]
	page := Page{Title: "data", Body: []byte(body)}
	page.save()
	http.Redirect(res, req, "/view/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/",http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8081", nil)
}

