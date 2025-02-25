package main

import (
	"backend/api/boxes"
	"backend/api/notes"
	notesboxes "backend/api/notes_boxes"
	"backend/api/users"
	database "backend/db/postgres"
	middleware "backend/middleware"
	model "backend/models/generated_model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	conn, connErr := database.NewPostgres()
	if connErr != nil {
		panic(connErr)
	}
	defer conn.DB.Close(context.Background())

	// Instantiate services
	queries := model.New(conn.DB)
	userService := users.NewService(queries)
	notesService := notes.NewService(queries)
	boxesService := boxes.NewService(queries)
	notesboxesService := notesboxes.NewService(queries)

	// Register handlers
	r := mux.NewRouter()
	userService.RegisterHandlers(r)
	notesService.RegisterHandlers(r)
	boxesService.RegisterHandlers(r)
	notesboxesService.RegisterHandlers(r)

	// Add middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	// Start the server
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
