package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"fdecono.com/internal/models"
)

type ProjectsHandler struct {
	projects []models.Project
}

func NewProjectsHandler(projects []models.Project) *ProjectsHandler {
	return &ProjectsHandler{projects: projects}
}

func (h *ProjectsHandler) List(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/projects.html",
	))

	data := struct {
		Title    string
		Projects []models.Project
	}{
		Title:    "Projects - fdecono",
		Projects: h.projects,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectsHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Extract project ID from URL path
	projectID := strings.TrimPrefix(r.URL.Path, "/projects/")
	if projectID == "" {
		http.NotFound(w, r)
		return
	}

	project := models.GetProjectByID(projectID)
	if project == nil {
		http.NotFound(w, r)
		return
	}

	// Check if this is an htmx request for partial content
	if r.Header.Get("HX-Request") == "true" {
		// Return just the project detail content
		tmpl := template.Must(template.ParseFiles("internal/templates/project-detail.html"))
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, project); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Full page request
	tmpl := template.Must(template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/project-detail.html",
	))

	data := struct {
		Title   string
		Project *models.Project
	}{
		Title:   project.Title + " - fdecono",
		Project: project,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
