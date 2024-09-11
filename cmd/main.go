package main

import (
	"log"
	"net/http"
	"online/internal/handlers"
	"online/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/survey/create", middleware.AuthMiddleware(handlers.CreateSurvey))
	mux.HandleFunc("/api/survey/vote", handlers.Vote)
	mux.HandleFunc("/api/survey/results", handlers.GetSurveyResults)

	mux.HandleFunc("/api/register", handlers.Register)
	mux.HandleFunc("/api/login", handlers.Login)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
