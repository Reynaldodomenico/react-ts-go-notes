package router

import (
	"net/http"

	"github.com/reynaldodomenico/react-and-go-notes/go-backend/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetNotes(w, r)
		case http.MethodPost:
			handlers.CreateNote(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/notes/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handlers.DeleteNote(w, r)
			return
		}
		http.NotFound(w, r)
	})
}
