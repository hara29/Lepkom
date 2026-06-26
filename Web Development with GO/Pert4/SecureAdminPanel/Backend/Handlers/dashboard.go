package Handlers

import (
	"html/template"
	"net/http"

	"SecureAdminPanel/Backend/config"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session-name")

	role, ok := session.Values["role"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if role == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	if role == "user" {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session-name")

	data := map[string]interface{}{
		"Username": session.Values["username"],
	}

	tmpl := template.Must(template.ParseFiles("Frontend/template/admin.html"))
	tmpl.Execute(w, data)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session-name")

	data := map[string]interface{}{
		"Username": session.Values["username"],
	}

	tmpl := template.Must(template.ParseFiles("Frontend/template/user.html"))
	tmpl.Execute(w, data)
}
