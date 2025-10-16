package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Job struct {
	ID       int    `json:"id"`
	Company  string `json:"company"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

var jobs = []Job{
	{ID: 1, Company: "Google", Role: "Backend Engineer", Status: "Applied", Notes: "Excited"},
}

// ✅ Simple Landing Page for Recruiters — IMPORTANT
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `
		<html>
		<head><title>JobTrack API</title></head>
		<body style="font-family: Arial; text-align:center; padding-top:50px;">
			<h1>🛠 JobTrack API</h1>
			<p>Status: <strong>Running ✅</strong></p>
			<p>Available endpoints:</p>
			<pre>/healthz  → Health Check
/jobs     → List Jobs (GET)
/jobs     → Add Job (POST)</pre>
			<p><a href="/healthz">Try Health Check</a></p>
		</body>
		</html>
	`)
}

// ✅ Health Route
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// ✅ GET /jobs – List jobs
func getJobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

// ✅ POST /jobs – Add job
func addJobHandler(w http.ResponseWriter, r *http.Request) {
	var newJob Job
	if err := json.NewDecoder(r.Body).Decode(&newJob); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newJob.ID = len(jobs) + 1
	jobs = append(jobs, newJob)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newJob)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			addJobHandler(w, r)
		} else {
			getJobsHandler(w, r)
		}
	})

	fmt.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
