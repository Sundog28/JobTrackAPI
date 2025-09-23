package main

import (
  "encoding/json"
  "log"
  "net/http"
  "sync"
)

type Job struct {
  ID int `json:"id"`
  Company string `json:"company"`
  Role string `json:"role"`
  Status string `json:"status"`
  Notes string `json:"notes"`
}

var (
  jobs []Job
  mu   sync.Mutex
  nextID = 1
)

func main() {
  http.HandleFunc("/jobs", handleJobs)
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){ w.WriteHeader(http.StatusOK) })
  log.Println("🚀 JobTrack API running on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleJobs(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  switch r.Method {
  case http.MethodGet:
    json.NewEncoder(w).Encode(jobs)
  case http.MethodPost:
    var j Job
    if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest); return
    }
    mu.Lock()
    j.ID = nextID; nextID++
    jobs = append(jobs, j)
    mu.Unlock()
    json.NewEncoder(w).Encode(j)
  default:
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
  }
}
