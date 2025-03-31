package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main(){

	dsn := "app_user:app_password@tcp(localhost:3306)/series_tracker?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	// router
	r:= chi.NewRouter()
	r.Use(middleware.Logger)

	http.ListenAndServe(":8080",r)
}