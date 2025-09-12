package models

import "time"

type Project struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	LongDesc    string    `json:"long_desc"`
	GitHubURL   string    `json:"github_url"`
	LiveURL     string    `json:"live_url,omitempty"`
	TechStack   []string  `json:"tech_stack"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Featured    bool      `json:"featured"`
}

func GetProjects() []Project {
	return []Project{
		{
			ID:          "go-htmx-website",
			Title:       "Personal Website",
			Description: "A modern personal website built with Go and htmx",
			LongDesc:    "This is my portfolio website. Built with Go for the backend, htmx for progressive enhancement, and Tailwind CSS for styling. Features include dynamic project listings, contact forms, and responsive design.",
			GitHubURL:   "https://github.com/fdecono/fdecono.com",
			// LiveURL:     "https://fdecono.com",
			TechStack: []string{"Go", "htmx"},
			CreatedAt: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			Featured:  true,
		},
		{
			ID:          "go-api-service",
			Title:       "REST API Service",
			Description: "A production-ready REST API template built with Ruby on Rails",
			LongDesc:    "A comprehensive Ruby on Rails API template for building RESTful APIs with authentication, testing, and documentation.",
			GitHubURL:   "https://github.com/fdecono/rails-api-template",
			TechStack:   []string{"Ruby on Rails", "PostgreSQL", "OAuth 2.0", "Swagger"},
			CreatedAt:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Featured:    true,
		},
	}
}

func GetProjectByID(id string) *Project {
	projects := GetProjects()
	for _, project := range projects {
		if project.ID == id {
			return &project
		}
	}
	return nil
}
