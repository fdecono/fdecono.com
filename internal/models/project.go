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
			LongDesc:    "This is a full-stack personal website showcasing my projects and skills. Built with Go for the backend, htmx for progressive enhancement, and Tailwind CSS for styling. Features include dynamic project listings, contact forms, and responsive design.",
			GitHubURL:   "https://github.com/fdecono.com/personal-website",
			LiveURL:     "https://fdecono.com",
			TechStack:   []string{"Go", "htmx", "Tailwind CSS", "HTML5"},
			// ImageURL:    "/static/images/fd.png",
			CreatedAt: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			Featured:  true,
		},
		{
			ID:          "go-api-service",
			Title:       "REST API Service",
			Description: "A production-ready REST API built with Go",
			LongDesc:    "A comprehensive REST API service with authentication, rate limiting, database integration, and comprehensive testing. Includes Docker configuration and CI/CD pipeline setup.",
			GitHubURL:   "https://github.com/fdecono.com/go-api-service",
			TechStack:   []string{"Go", "PostgreSQL", "Docker", "JWT", "Redis"},
			CreatedAt:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Featured:    true,
		},
		{
			ID:          "react-dashboard",
			Title:       "React Dashboard",
			Description: "A modern admin dashboard built with React and TypeScript",
			LongDesc:    "A feature-rich admin dashboard with data visualization, user management, and real-time updates. Built with modern React patterns and TypeScript for type safety.",
			GitHubURL:   "https://github.com/fdecono.com/react-dashboard",
			LiveURL:     "https://dashboard-demo.fdecono.com",
			TechStack:   []string{"React", "TypeScript", "Chart.js", "Material-UI"},
			CreatedAt:   time.Date(2023, 12, 20, 0, 0, 0, 0, time.UTC),
			Featured:    false,
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
