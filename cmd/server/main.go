package main

import (
	"log"
	"net/http"
	"os"

	"fdecono.com/internal/handlers"
	"fdecono.com/internal/models"
)

func main() {
	// Initialize project data
	projects := models.GetProjects()

	// Create handlers
	homeHandler := handlers.NewHomeHandler(projects)
	projectsHandler := handlers.NewProjectsHandler(projects)
	contactHandler := handlers.NewContactHandler()

	// Setup routes
	mux := http.NewServeMux()

	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	// Main routes
	mux.HandleFunc("/", homeHandler.Home)
	mux.HandleFunc("/projects", projectsHandler.List)
	mux.HandleFunc("/projects/", projectsHandler.Detail)
	mux.HandleFunc("/contact", contactHandler.Contact)
	mux.HandleFunc("/contact/submit", contactHandler.Submit)

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
