package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1 dbname=bioskop sslmode=disable"
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		lokasi VARCHAR(100) NOT NULL,
		rating FLOAT
	);`

	DB.MustExec(schema)
	fmt.Println("Database siap digunakan 🚀")
}
