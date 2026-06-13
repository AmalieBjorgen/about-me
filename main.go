package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amaliebjorgen/about-me/api"
)

func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.Handle("/images/", fs)

	// Route all other pages and endpoint requests to the Vercel handler
	http.HandleFunc("/", handler.Handler)

	// Read Vercel's assigned PORT environment variable, fallback to 8080 locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("🚀 Server running on %s", addr)
	log.Printf("👉 Press Ctrl+C to stop")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
