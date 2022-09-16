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
	// >> static files start
	MainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// mux path route
	// styles
	MainRouter.Handle("/{something}/css", http.StripPrefix("/css", http.FileServer(http.Dir("static/css"))))
	MainRouter.Handle("/css/{style}", http.StripPrefix("/css", http.FileServer(http.Dir("static/css"))))
	// scripts
	MainRouter.Handle("/js/vendor/{script}", http.StripPrefix("/js/vendor", http.FileServer(http.Dir("static/js/vendor"))))
	MainRouter.Handle("/js/{script}", http.StripPrefix("/js", http.FileServer(http.Dir("static/js"))))
	// imgs
	MainRouter.Handle("/img/icon/{image}", http.StripPrefix("/img/icon", http.FileServer(http.Dir("static/img/icon"))))
	MainRouter.Handle("/img/svg_icon/{image}", http.StripPrefix("/img/svg_icon", http.FileServer(http.Dir("static/img/svg_icon"))))
	//
	MainRouter.Handle("/img/about/{image}", http.StripPrefix("/img/about", http.FileServer(http.Dir("static/img/about"))))
	MainRouter.Handle("/img/banner/{image}", http.StripPrefix("/img/banner", http.FileServer(http.Dir("static/img/banner"))))
	MainRouter.Handle("/img/blog/{image}", http.StripPrefix("/img/blog", http.FileServer(http.Dir("static/img/blog"))))
	MainRouter.Handle("/img/comment/{image}", http.StripPrefix("/img/comment", http.FileServer(http.Dir("static/img/comment"))))
	MainRouter.Handle("/img/delicious/{image}", http.StripPrefix("/img/delicious", http.FileServer(http.Dir("static/img/delicious"))))
	MainRouter.Handle("/img/elements/{image}", http.StripPrefix("/img/elements", http.FileServer(http.Dir("static/img/elements"))))
	MainRouter.Handle("/img/gallery/{image}", http.StripPrefix("/img/gallery", http.FileServer(http.Dir("static/img/gallery"))))
	MainRouter.Handle("/img/post/{image}", http.StripPrefix("/img/post", http.FileServer(http.Dir("static/img/post"))))
	MainRouter.Handle("/img/service/{image}", http.StripPrefix("/img/service", http.FileServer(http.Dir("static/img/service"))))
	MainRouter.Handle("/img/testimonial/{image}", http.StripPrefix("/img/testimonial", http.FileServer(http.Dir("static/img/testimonial"))))
	//MainRouter.Handle("/img//{image}", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img/"))))
	// images
	MainRouter.Handle("/images/{image}", http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))
	MainRouter.Handle("/img/{image}", http.StripPrefix("/img", http.FileServer(http.Dir("static/img"))))
	// fonts and icons
	MainRouter.Handle("/fonts/{font}", http.StripPrefix("/fonts", http.FileServer(http.Dir("static/fonts"))))
	MainRouter.Handle("/icon/{icon}", http.StripPrefix("/icon", http.FileServer(http.Dir("static/icon"))))
	// custom images
	MainRouter.Handle("/static/images/products/{file}", http.StripPrefix("/static/images/products", http.FileServer(http.Dir("static/images/products"))))
	// >> static files end

	// test
	//testSubRouter := MainRouter.PathPrefix("/test").Subrouter()
	//testSubRouter.HandleFunc("", testRoute)
	//testSubRouter.HandleFunc("/", testRoute)

	// index
	IndexSubRouter := MainRouter.PathPrefix("").Subrouter()
	IndexSubRouter.HandleFunc("/", IndexPage)
	// about
	AboutSubRouter := MainRouter.PathPrefix("/about").Subrouter()
	AboutSubRouter.HandleFunc("", AboutPage)
	AboutSubRouter.HandleFunc("/", AboutPage)
	// menu
	MenuSubRouter := MainRouter.PathPrefix("/menu").Subrouter()
	MenuSubRouter.HandleFunc("", MenuPage)
	MenuSubRouter.HandleFunc("/", MenuPage)
	// blog
	BlogSubRouter := MainRouter.PathPrefix("/blog").Subrouter()
	BlogSubRouter.HandleFunc("", BlogPage)
	BlogSubRouter.HandleFunc("/", BlogPage)
	// contact
	ContactSubRouter := MainRouter.PathPrefix("/contact").Subrouter()
	ContactSubRouter.HandleFunc("", ContactPage)
	ContactSubRouter.HandleFunc("/", ContactPage)
	// elements
	ElementsSubRouter := MainRouter.PathPrefix("/elements").Subrouter()
	ElementsSubRouter.HandleFunc("", ElementsPage)
	ElementsSubRouter.HandleFunc("/", ElementsPage)
	// single-blog
	Single_blogSubRouter := MainRouter.PathPrefix("/single_blog").Subrouter()
	Single_blogSubRouter.HandleFunc("", Single_blogPage)
	Single_blogSubRouter.HandleFunc("/", Single_blogPage)
	// return main router
	return MainRouter
}

func RouterReady() *mux.Router {
	MainRouter := RouterUpdate()
	return MainRouter
}
