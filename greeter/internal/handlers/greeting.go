package handlers

import (
	"context"
	"fmt"

	"greeter/internal/gen"
)

// Server implements gen.StrictServerInterface.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetGreeting(ctx context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name
	if name == "" {
		name = "World"
	}
	return gen.GetGreeting200JSONResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
