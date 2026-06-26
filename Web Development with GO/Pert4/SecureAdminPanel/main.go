package main

import (
	"log"
	"net/http"

	"SecureAdminPanel/Backend/Handlers"
	"SecureAdminPanel/Backend/Middleware"
	"SecureAdminPanel/Backend/config"
)

func main() {
	config.InitDB()

	fs := http.FileServer(http.Dir("Frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.HandleFunc("/login", Handlers.LoginPage)
	http.HandleFunc("/login-process", Handlers.LoginHandler)
	http.HandleFunc("/logout", Handlers.LogoutHandler)

	http.Handle("/dashboard",
		Middleware.IsAuthenticated(
			http.HandlerFunc(Handlers.DashboardHandler),
		),
	)

	http.Handle("/admin",
		Middleware.IsAuthenticated(
			Middleware.RoleOnly("admin",
				http.HandlerFunc(Handlers.AdminHandler),
			),
		),
	)

	http.Handle("/user",
		Middleware.IsAuthenticated(
			Middleware.RoleOnly("user",
				http.HandlerFunc(Handlers.UserHandler),
			),
		),
	)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
