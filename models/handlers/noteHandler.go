package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/yourusername/notes-api/models"
    "math/rand"
    "strconv"
)

var notes []models.Note

// Get all notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(notes)
}

// Get single note
func GetNote(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range notes {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Note not found", http.StatusNotFound)
}

// Create new note
func CreateNote(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var note models.Note
    _ = json.NewDecoder(r.Body).Decode(&note)
    note.ID = strconv.Itoa(rand.Intn(100000)) // simple random ID
    notes = append(notes, note)
    json.NewEncoder(w).Encode(note)
}

// Delete note
func DeleteNote(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range notes {
        if item.ID == params["id"] {
            notes = append(notes[:index], notes[index+1:]...)
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Note not found", http.StatusNotFound)
}
