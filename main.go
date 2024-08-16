package main

import (
	"database/sql"

	"github.com/pedroesteves2803/arquitetura-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/pedroesteves2803/arquitetura-hexagonal/adapters/db"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	productService.Create("Produto xemplo", 20)
}
