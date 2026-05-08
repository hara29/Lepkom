package main

import (
	"fmt"
	"net/http"

	"toko_lepkom_npm/db"
	"toko_lepkom_npm/handlers"
)

func main() {

	port := ":8080"

	db.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeView)
	mux.HandleFunc("/image", handlers.ImageView)
	mux.HandleFunc("/edit", handlers.EditView)
	mux.HandleFunc("/update", handlers.UpdateProductHandler)
	mux.HandleFunc("/create", handlers.CreateProductHandler)
	mux.HandleFunc("/delete", handlers.DeleteProductHandler)

	fmt.Printf("Server running at http://localhost%s\n", port)

	http.ListenAndServe(port, mux)
}
