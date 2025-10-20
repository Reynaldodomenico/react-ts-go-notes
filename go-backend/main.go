package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var (
	notes  = []Note{}
	nextID = 1
	notesMu sync.Mutex
)

func main() {
	http.HandleFunc("/api/notes", notesHandler)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, corsMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w, r)
	case http.MethodPost:
		createNote(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	notesMu.Lock()
	defer notesMu.Unlock()
	json.NewEncoder(w).Encode(notes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in struct{ Text string `json:"text"` }
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
}

notesMu.Lock()
	defer notesMu.Unlock()
	n := Note{ID: nextID, Text: in.Text}
	nextID++
	notes = append(notes, n)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}