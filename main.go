package main

import (
	"log"
	"net/http"

	"github.com/amaliebjorgen/about-me/api"
)

func main() {
	// Serve static files from the "public" directory locally
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.Handle("/images/", fs)

	// Route all other pages and endpoint requests to the Vercel handler
	http.HandleFunc("/", handler.Handler)

	addr := ":8080"
	log.Printf("🚀 Local development server running at http://localhost%s", addr)
	log.Printf("👉 Press Ctrl+C to stop")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
