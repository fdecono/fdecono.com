package handlers

import (
	"html/template"
	"net/http"

	"fdecono.com/internal/models"
)

type HomeHandler struct {
	projects []models.Project
}

func NewHomeHandler(projects []models.Project) *HomeHandler {
	return &HomeHandler{projects: projects}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/home.html",
	))

	data := struct {
		Title    string
		Projects []models.Project
	}{
		Title:    "Federico Decono",
		Projects: h.projects[:2],
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
