package handlers

import (
	"html/template"
	"net/http"
	"strings"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

type ContactForm struct {
	Name    string
	Email   string
	Subject string
	Message string
	Errors  map[string]string
}

func (h *ContactHandler) Contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/contact.html",
	))

	data := struct {
		Title string
		Form  ContactForm
	}{
		Title: "Contact - fdecono",
		Form:  ContactForm{},
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ContactHandler) Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	form := ContactForm{
		Name:    strings.TrimSpace(r.FormValue("name")),
		Email:   strings.TrimSpace(r.FormValue("email")),
		Subject: strings.TrimSpace(r.FormValue("subject")),
		Message: strings.TrimSpace(r.FormValue("message")),
		Errors:  make(map[string]string),
	}

	// Validate form
	if form.Name == "" {
		form.Errors["name"] = "Name is required"
	}
	if form.Email == "" {
		form.Errors["email"] = "Email is required"
	} else if !strings.Contains(form.Email, "@") {
		form.Errors["email"] = "Please enter a valid email"
	}
	if form.Subject == "" {
		form.Errors["subject"] = "Subject is required"
	}
	if form.Message == "" {
		form.Errors["message"] = "Message is required"
	}

	// Check if this is an htmx request
	if r.Header.Get("HX-Request") == "true" {
		if len(form.Errors) > 0 {
			// Return form with errors
			tmpl := template.Must(template.ParseFiles("internal/templates/contact-form.html"))
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, form)
			return
		}

		// Send email (in production, you'd use a proper email service)
		if err := h.sendEmail(form); err != nil {
			form.Errors["general"] = "Failed to send message. Please try again."
			tmpl := template.Must(template.ParseFiles("internal/templates/contact-form.html"))
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, form)
			return
		}

		// Success response
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<div class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded mb-4">
				<strong>Success!</strong> Your message has been sent. I'll get back to you soon.
			</div>
		`))
		return
	}

	// Regular form submission
	if len(form.Errors) > 0 {
		tmpl := template.Must(template.ParseFiles(
			"internal/templates/base.html",
			"internal/templates/contact.html",
		))
		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "base", struct {
			Title string
			Form  ContactForm
		}{Title: "Contact - fdecono", Form: form})
		return
	}

	// Send email and redirect
	if err := h.sendEmail(form); err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contact?success=true", http.StatusSeeOther)
}

func (h *ContactHandler) sendEmail(form ContactForm) error {
	// In production, you'd use a proper email service like SendGrid, AWS SES, etc.
	// For now, we'll just log the message
	// You can implement actual email sending here
	
	// Example using SMTP (you'd need to configure this)
	// smtpHost := os.Getenv("SMTP_HOST")
	// smtpPort := os.Getenv("SMTP_PORT")
	// smtpUser := os.Getenv("SMTP_USER")
	// smtpPass := os.Getenv("SMTP_PASS")
	
	// For demo purposes, just log the message
	println("Contact form submission:")
	println("Name:", form.Name)
	println("Email:", form.Email)
	println("Subject:", form.Subject)
	println("Message:", form.Message)
	
	return nil
}
