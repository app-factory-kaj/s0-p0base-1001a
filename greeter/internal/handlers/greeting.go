package handlers

import (
	"context"
	"fmt"

	"greeter/internal/gen"
)

type GreetingServer struct{}

func (GreetingServer) GetGreeting(_ context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name
	if name == "" {
		name = "World"
	}
	return gen.GetGreeting200JSONResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
