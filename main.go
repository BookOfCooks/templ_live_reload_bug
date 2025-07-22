package main

import (
	"fmt"
	"main/templates"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Home().Render(r.Context(), w)
	})

	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
