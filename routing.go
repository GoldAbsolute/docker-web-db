package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/about.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
func MenuPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/menu.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}

// <<<--- TESTING DB --->>>
func testRoute(w http.ResponseWriter, r *http.Request) {

	var app_login = os.Getenv("app_login")
	var app_password = os.Getenv("app_password")
	var app_ip = os.Getenv("app_ip")
	var app_port = os.Getenv("app_port")
	var app_dbname = os.Getenv("app_dbname")

	var db_path = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", app_login, app_password, app_ip, app_port, app_dbname)

	db, _ := sql.Open("mysql", db_path)
	//db, err := sql.Open("mysql", "root:db_pass@(127.0.0.1:3306)/my-db?parseTime=true")
	//check(err)
	//_ = db
	_ = db.Ping()
	data := "1"
	_ = data
	_ = r
	tmpl := template.Must(template.ParseFiles("tmpl/template.html", "tmpl/header.html"))
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}
