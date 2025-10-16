package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)       // Friendly landing page
	http.HandleFunc("/healthz", handleHealth) // Health check endpoint

	// Existing backend routes (keep your /jobs CRUD etc.)
	// Example:
	// http.HandleFunc("/jobs", handleJobs)

	fmt.Println("🚀 JobTrack API is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Landing page for recruiters or visitors
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>JobTrack API</title>
			<style>
				body { font-family: Arial, sans-serif; background: #0d1117; color: #c9d1d9; text-align: center; padding-top: 60px; }
				.container { max-width: 600px; margin: auto; }
				.code { background: #161b22; padding: 10px; border-radius: 6px; margin-top: 20px; }
				a { color: #58a6ff; text-decoration: none; }
			</style>
		</head>
		<body>
			<div class="container">
				<h1>🛠 JobTrack API</h1>
				<p>Status: <strong style="color: #2ea043;">Running ✅</strong></p>
				<p>Use <code>/healthz</code> to check API health.</p>
				<p>Use <code>/jobs</code> (if implemented) to manage job entries.</p>
				<br>
				<em>Deployed via Render • Built by John Treen</em>
			</div>
		</body>
		</html>
	`)
}

// Health check
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
