package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	r := chi.NewRouter()
	srv := handlers.NewServer()
	handler := gen.HandlerWithOptions(gen.NewStrictHandler(srv, nil), gen.ChiServerOptions{BaseRouter: r})

	log.Printf("greeter listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
