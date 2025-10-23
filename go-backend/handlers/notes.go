package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/db"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/models"
)

func GetNotes(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, text, created_at FROM notes ORDER BY id DESC")
    if err != nil {
        log.Println("❌ DB query error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var notes []models.Note
    for rows.Next() {
        var n models.Note
        if err := rows.Scan(&n.ID, &n.Text, &n.CreatedAt); err != nil {
            log.Println("❌ Row scan error:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        notes = append(notes, n)
    }

    c.JSON(http.StatusOK, notes)
}


func CreateNote(c *gin.Context) {
	var n models.Note
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.DB.QueryRow(
		`INSERT INTO notes (text) VALUES ($1) RETURNING id, text, created_at`,
		n.Text,
	).Scan(&n.ID, &n.Text, &n.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, n)
}

func DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = db.DB.Exec(`DELETE FROM notes WHERE id=$1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
