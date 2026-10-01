package main

import (
	"log"
	"net/http"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	srv := handlers.GreetingServer{}
	handler := gen.HandlerWithOptions(gen.NewStrictHandler(srv, nil), gen.ChiServerOptions{})

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatal(err)
	}
}
