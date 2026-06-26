package Handlers

import (
	"html/template"
	"net/http"

	"SecureAdminPanel/Backend/config"

	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Frontend/template/login.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, exists := config.GetUserByUsername(username)
	if !exists {
		http.Error(w, "Username atau password salah!", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Username atau password salah!", http.StatusUnauthorized)
		return
	}

	session, _ := config.Store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role
	session.Save(r, w)

	if user.Role == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session-name")

	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
