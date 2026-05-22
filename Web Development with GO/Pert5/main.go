package main

import (
	"log"
	"net/http"
	"pert5/handlers"
	"pert5/middlewares"
)

func main() {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTasks(w, r)
		case http.MethodPost:
			handlers.CreateTask(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTaskByID(w, r)
		case http.MethodPut:
			handlers.UpdateTask(w, r)
		case http.MethodPatch:
			handlers.PatchTask(w, r)
		case http.MethodDelete:
			handlers.DeleteTask(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	// Serve static files (frontend HTML, CSS, JS)
	// Misalnya index.html ada di folder "static"
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", middlewares.Logger(mux))
}
