package models

import (
	"strings"
	"io/ioutil"
	"../../config"
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

func (p *Page) Save() error {
	filename := config.DATA_PATH + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (StringPage, error) {
	filename := config.DATA_PATH + title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	page := Page{Title: title, Body: body}
	return parsePage(&page), nil
}

func ListPage() []string{
	files, _ := ioutil.ReadDir(config.DATA_PATH)
	names := []string{}
	for i := 0;i < len(files); i++  {
		if(files[i].IsDir()) { continue }

		name := strings.TrimSuffix(files[i].Name(),".txt")
		names = append(names, name)
	}

	return names
}