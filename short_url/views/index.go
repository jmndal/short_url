package views

import (
	"fmt"
	"net/http"
	"text/template"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	c := map[string]interface{}{
		"Title": "Short URL",
	}

	tmpl.Execute(w, c)
}

func GormDB() *gorm.DB {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/short_url?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}
