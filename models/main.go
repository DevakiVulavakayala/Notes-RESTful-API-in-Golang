package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/yourusername/notes-api/handlers"
)

func main() {
    router := mux.NewRouter()

    // Sample note
    handlers.Notes = append(handlers.Notes, models.Note{ID: "1", Title: "Sample", Content: "This is a sample note"})

    // Route handlers
    router.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
    router.HandleFunc("/notes/{id}", handlers.GetNote).Methods("GET")
    router.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
    router.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

    log.Println("Server started at :8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
