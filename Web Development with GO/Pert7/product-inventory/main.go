package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"product-inventory/configs"
	"product-inventory/handlers"
	"product-inventory/middlewares"
	"product-inventory/utils"
)

func main() {
	configs.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RedirectHome)
	mux.HandleFunc("/login", handlers.ServeLogin)
	mux.HandleFunc("/products", handlers.ServeProductsPage)
	mux.HandleFunc("/users", handlers.ServeUsersPage)
	mux.HandleFunc("/api/login", method(http.MethodPost, handlers.Login))

	mux.Handle("/api/register", middlewares.AdminOnly(http.HandlerFunc(method(http.MethodPost, handlers.Register))))
	mux.Handle("/api/users", middlewares.AdminOnly(http.HandlerFunc(method(http.MethodGet, handlers.ListUsers))))
	mux.Handle("/api/generate-token", middlewares.AdminOnly(http.HandlerFunc(method(http.MethodPost, handlers.GenerateAPIToken))))
	mux.Handle("/api/user/token", middlewares.AdminOnly(http.HandlerFunc(method(http.MethodPost, handlers.GenerateAPIToken))))

	mux.Handle("/api/products", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProducts(w, r)
		case http.MethodPost:
			handlers.CreateProduct(w, r)
		default:
			utils.WriteJSON(w, http.StatusMethodNotAllowed, "error", "Method tidak diizinkan", nil)
		}
	})))

	mux.Handle("/api/products/", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/products/") {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handlers.GetProductByID(w, r)
		case http.MethodPut:
			handlers.UpdateProduct(w, r)
		case http.MethodDelete:
			handlers.DeleteProduct(w, r)
		default:
			utils.WriteJSON(w, http.StatusMethodNotAllowed, "error", "Method tidak diizinkan", nil)
		}
	})))

	mux.Handle("/api/transactions", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTransactions(w, r)
		case http.MethodPost:
			handlers.CreateTransaction(w, r)
		default:
			utils.WriteJSON(w, http.StatusMethodNotAllowed, "error", "Method tidak diizinkan", nil)
		}
	})))

	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, middlewares.Logger(mux)))
}

func method(allowed string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowed {
			utils.WriteJSON(w, http.StatusMethodNotAllowed, "error", "Method tidak diizinkan", nil)
			return
		}
		handler(w, r)
	}
}
