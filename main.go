package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jmandal/short_url/api"
	"github.com/jmandal/short_url/models"
	"github.com/jmandal/short_url/views"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "localhost"
	Port   = ":1666"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("short_url")
	MigrateDB()
	Handlers()

	http.ListenAndServe(Port, nil)
}

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", api.RedirectHandler)
	http.HandleFunc("/short-url/", views.IndexHandler)
	http.HandleFunc("/shorten/", api.ShortenHandler)
}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	url_short_list := models.URLShortList{}

	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/short_url?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&url_short_list)
}
