package handlers

import (
	"context"
	"strings"
	"testing"

	"greeter/internal/gen"
)

func TestGetGreetingWithName(t *testing.T) {
	s := NewServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: "Alice"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	greeting, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type %T", resp)
	}
	if !strings.Contains(greeting.Message, "Alice") {
		t.Errorf("expected message to contain %q, got %q", "Alice", greeting.Message)
	}
}

func TestGetGreetingWithoutName(t *testing.T) {
	s := NewServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: ""},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	greeting, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type %T", resp)
	}
	if !strings.Contains(greeting.Message, "World") {
		t.Errorf("expected message to contain %q, got %q", "World", greeting.Message)
	}
}

func TestGetGreetingWithEmptyName(t *testing.T) {
	s := NewServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: ""},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	greeting, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type %T", resp)
	}
	if !strings.Contains(greeting.Message, "World") {
		t.Errorf("expected message to contain %q, got %q", "World", greeting.Message)
	}
}
