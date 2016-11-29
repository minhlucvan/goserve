package main

import (
	"net/http"
	"./config/"
	"./server/router"
	"./server/controllers"
)

func setRouters(){
	router.Route("/", controllers.IndexHander)
	router.Route("/add/", controllers.AddHandler)
	router.Route("/edit/", controllers.EditHandler)
	router.Route("/save/", controllers.SaveHandler)
	router.Static("/assets/", config.ASSET_PATH)
}

func main() {
	setRouters()
	http.ListenAndServe(":8081", nil)
}

