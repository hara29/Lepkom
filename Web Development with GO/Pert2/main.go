package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title         string
	Company       string
	TotalProjects int
	Description   string
	Image         string
	Projects      []Project
}
type Project struct {
	Name        string
	Description string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles(
		"templates/layout.html",
		"templates/"+tmpl,
	)

	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "layout", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	data := PageData{
		Title:   "Home",
		Company: "LePKom",
	}

	renderTemplate(w, "home.html", data)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title:       "Profile",
		Company:     "LePKom",
		Description: "Lembaga kursus IT untuk persiapan uji kompetensi",
		Image:       "/img/logo.png",
	}

	renderTemplate(w, "profile.html", data)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {

	projects := []Project{
		{Name: "Website", Description: "Development of a responsive website"},
		{Name: "Mobile App", Description: "Creation of a cross-platform mobile application"},
		{Name: "API System", Description: "Design and implementation of a RESTful API"},
	}

	data := PageData{
		Title:         "Projects",
		Company:       "LePKom",
		TotalProjects: len(projects),
		Projects:      projects,
	}

	renderTemplate(w, "projects.html", data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	data := PageData{
		Title: "404 Not Found",
	}

	renderTemplate(w, "404.html", data)
}

func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := ":8088"
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/profile", ProfileHandler)
	mux.HandleFunc("/projects", ProjectsHandler)

	fs := http.FileServer(http.Dir("./img"))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))

	handler := loggingMiddleware(mux)

	fmt.Printf("Server is running on http://localhost%s", port)
	http.ListenAndServe(port, handler)
}
