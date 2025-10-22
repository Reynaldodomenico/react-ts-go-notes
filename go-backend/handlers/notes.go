package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/reynaldodomenico/react-and-go-notes/go-backend/db"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/models"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT id, text, created_at FROM notes ORDER BY id DESC`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Text, &n.CreatedAt); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		notes = append(notes, n)
	}
	json.NewEncoder(w).Encode(notes)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}

	err := db.DB.QueryRow(
		`INSERT INTO notes (text) VALUES ($1) RETURNING id, text, created_at`,
		n.Text,
	).Scan(&n.ID, &n.Text, &n.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/notes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}
	_, err = db.DB.Exec(`DELETE FROM notes WHERE id=$1`, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
