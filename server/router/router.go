package router

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type Routes map[string]HandlerFunc

var roures = Routes{}

func Route(path string, handler HandlerFunc){
	http.HandleFunc(path, handler)
}

func Static(path string, dir string ){
	fs := http.FileServer(http.Dir(dir))
	http.Handle(path, http.StripPrefix(path, fs))
}
