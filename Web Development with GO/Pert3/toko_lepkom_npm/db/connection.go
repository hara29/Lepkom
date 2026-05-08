package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB membuka koneksi ke database MySQL dan menyimpannya ke variabel global DB.
func ConnectDB() {
	// DSN berisi konfigurasi user, password, host, nama database, dan opsi parsing waktu.
	dsn := "root:@tcp(localhost:3306)/toko_lepkom_npm?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ping digunakan untuk memastikan database benar-benar dapat diakses.
	err = db.Ping()
	if err != nil {
		log.Fatal("Database not connected:", err)
	}

	fmt.Println("Database connected successfully!")
	DB = db
}
