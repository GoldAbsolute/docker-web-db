package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RouterCreate() *mux.Router {
	MainRouter := mux.NewRouter()
	return MainRouter
}

func RouterUpdate() *mux.Router {
	// create
	MainRouter := RouterCreate()
	// static files start
	MainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// static files end

	// routes
	//MainRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	tmpl := template.Must(template.ParseFiles("tmpl/template.html", "tmpl/header.html"))
	//	errT := tmpl.ExecuteTemplate(writer, "index", nil)
	//	if errT != nil {
	//		panic(errT)
	//	}
	//})

	//testSubRouter := MainRouter.PathPrefix("/test").Subrouter()
	//testSubRouter.HandleFunc("", testRoute)
	//testSubRouter.HandleFunc("/", testRoute)
	// index
	IndexSubRouter := MainRouter.PathPrefix("/").Subrouter()
	IndexSubRouter.HandleFunc("/", IndexPage)
	// about
	AboutSubRouter := MainRouter.PathPrefix("/about").Subrouter()
	AboutSubRouter.HandleFunc("/", AboutPage)
	// menu
	MenuSubRouter := MainRouter.PathPrefix("/menu").Subrouter()
	MenuSubRouter.HandleFunc("/", MenuPage)

	return MainRouter
}

func RouterReady() *mux.Router {
	MainRouter := RouterUpdate()
	return MainRouter
}
