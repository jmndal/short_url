package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/jmandal/short_url/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "localhost"
	Port   = ":1666"
)

// RedirectHandler handles requests for shortened URLs and redirects to the original URL.
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	// Extract the short token from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/short-url/")
	id = strings.TrimPrefix(id, "/")

	fmt.Printf("Shortened 2222222222: %s\n", id)

	var url models.URLShortList
	db.First(&url, "shorten = ?", id)

	fmt.Printf("Redirecting to Original URL: %s\n", url.Original)
	http.Redirect(w, r, url.Original, http.StatusFound)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data to get the long URL
	r.ParseForm()
	longURL := r.Form.Get("url")
	fmt.Println("Long URL: ", longURL)

	// Query existing records from the database
	urlShort := []models.URLShortList{}
	models.GormDB().Find(&urlShort)

	shorten := GenerateShortToken(5)

	// Check if the generated token already exists in the database
	for i := range urlShort {
		if urlShort[i].Shorten == shorten {
			// If the token exists, regenerate and restart the loop
			shorten = GenerateShortToken(5)
			i = -1 // Restart the loop
		}
	}

	// Save the new entry in the database
	urlList := models.URLShortList{}
	urlList.Original = longURL
	urlList.Shorten = shorten
	urlList.Save()

	shortURL := "http://" + BindIP + Port + "/" + shorten

	// Return the short URL in the response
	fmt.Fprintf(w, "%s", shortURL)
}

func GenerateShortToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
	key := make([]byte, length)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return string(key)
}

func GormDB() *gorm.DB {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/short_url?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}
