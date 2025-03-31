package db

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Cambia la dirección a la del contenedor MySQL
	dsn := "root:password@tcp(db:3306)/seriesdb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexión a la base de datos establecida.")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal(err)
	}
}
