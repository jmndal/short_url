package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type URLShortList struct {
	ID       uint `gorm:"primaryKey"`
	Original string
	Shorten  string
}

func (u *URLShortList) String() string {
	return u.Shorten
}

func (u *URLShortList) Save() {
	db := GormDB()

	db.Save(u)
}

func GormDB() *gorm.DB {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/short_url?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}
