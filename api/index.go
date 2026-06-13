package handler

import (
	"log"
	"net/http"

	"github.com/amaliebjorgen/about-me/api/models"
	"github.com/amaliebjorgen/about-me/api/templates"
)

var CVData = models.CVData

// Handler is the entrypoint for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set basic headers
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	isHTMX := r.Header.Get("HX-Request") == "true"
	path := r.URL.Path

	switch path {
	case "/", "/about":
		if isHTMX {
			err := templates.AboutContent(CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering About content: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			err := templates.Layout(CVData.Profile.Name+" | About Me", "about", templates.AboutContent(CVData), CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Layout: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

	case "/experience":
		if isHTMX {
			err := templates.ExperienceContent(CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Experience content: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			err := templates.Layout(CVData.Profile.Name+" | Experience", "experience", templates.ExperienceContent(CVData), CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Layout: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

	case "/projects":
		tag := r.URL.Query().Get("tag")
		if tag == "" {
			tag = "All"
		}

		target := r.Header.Get("HX-Target")
		if isHTMX && target == "project-cards-container" {
			err := templates.ProjectCardsList(CVData.Projects, tag).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Project cards list: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else if isHTMX {
			err := templates.ProjectsContent(CVData, tag).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Projects content: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			err := templates.Layout(CVData.Profile.Name+" | Projects", "projects", templates.ProjectsContent(CVData, tag), CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Layout: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

	case "/contact":
		if r.Method == http.MethodPost {
			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			log.Printf("New contact form submission: Name=%s, Email=%s, Msg=%s", name, email, message)

			// Render the success fragment
			err = templates.ContactSuccess(name).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Contact Success: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		if isHTMX {
			err := templates.ContactContent().Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Contact content: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			err := templates.Layout(CVData.Profile.Name+" | Contact", "contact", templates.ContactContent(), CVData).Render(r.Context(), w)
			if err != nil {
				log.Printf("Error rendering Layout: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

	default:
		// Let Vercel default file routing or local files handle other paths
		http.NotFound(w, r)
	}
}
