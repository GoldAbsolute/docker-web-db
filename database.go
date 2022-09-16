package main

import (
	"database/sql"
	"fmt"
	"os"
)

func ReturnDB() *sql.DB {
	//var app_login = os.Getenv("app_login")
	//var app_password = os.Getenv("app_password")
	//var app_ip = os.Getenv("app_ip")
	//var app_port = os.Getenv("app_port")
	//var app_dbname = os.Getenv("app_dbname")

	app_login, _ := os.LookupEnv("app_login")
	app_password, _ := os.LookupEnv("app_password")
	app_ip, _ := os.LookupEnv("app_ip")
	app_port, _ := os.LookupEnv("app_port")
	app_dbname, _ := os.LookupEnv("app_dbname")

	var db_path = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", app_login, app_password, app_ip, app_port, app_dbname)

	db, _ := sql.Open("mysql", db_path)

	//Решил добавить отсебятину
	//defer db.Close()

	return db
}

// TODO Это я вообще не сделал,настрой что возвращает и дб
func IndexInfo() IndexInfoObj {
	db := ReturnDB()
	rows, err := db.Query(`SELECT id, title, text, created_at FROM food_list ORDER BY id DESC;`)
	defer rows.Close()
	check(err)

	i := IndexInfoObj{}
	return i
}
