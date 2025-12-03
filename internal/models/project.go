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
			ID:          "imposter",
			Title:       "Imposter",
			Description: "A multiplayer word game where players find the imposter among them.",
			LongDesc:    "Imposter is a real-time multiplayer party game built with Go and WebSockets. Players take turns submitting words related to a secret word, but one player - the imposter - doesn't know what the word is. Can you blend in or catch the fake?",
			GitHubURL:   "https://imposter.fdecono.com",
			LiveURL:     "https://imposter.fdecono.com",
			TechStack:   []string{"Go", "WebSockets", "Real-time"},
			CreatedAt:   time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
			Featured:    true,
		},
		{
			ID:          "the-knight-solver",
			Title:       "The Knight Solver",
			Description: "Complete the Knight's Tour puzzle on a chessboard.",
			LongDesc:    "This is a implementation of the Knight's Tour puzzle. It uses a backtracking algorithm to find a solution to the puzzle.",
			GitHubURL:   "https://knight.fdecono.com",
			TechStack:   []string{"Go"},
			CreatedAt:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Featured:    true,
		},
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
			ID:          "go-mini-blockchain",
			Title:       "Mini Blockchain",
			Description: "A mini blockchain built with Go.",
			LongDesc:    "This implementation serves as a hands-on learning exercise to understand blockchain fundamentals.",
			GitHubURL:   "https://github.com/fdecono/mini_blockchain",
			TechStack:   []string{"Go", "blockchain", "p2p"},
			CreatedAt:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Featured:    true,
		},
		{
			ID:          "my-bit-torrent",
			Title:       "MyBitTorrent (WIP)",
			Description: "A BitTorrent client built with Go",
			LongDesc:    "A BitTorrent client built with Go. It can download and upload files to the BitTorrent network.",
			GitHubURL:   "https://github.com/fdecono/my-bit-torrent",
			TechStack:   []string{"Go", "p2p", "bittorrent"},
			CreatedAt:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			Featured:    true,
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
