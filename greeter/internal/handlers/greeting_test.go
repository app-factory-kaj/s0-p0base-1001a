package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func newHandler() http.Handler {
	return gen.HandlerWithOptions(gen.NewStrictHandler(handlers.GreetingServer{}, nil), gen.ChiServerOptions{})
}

func TestGetGreetingWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Alice", nil)
	rec := httptest.NewRecorder()
	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body struct{ Message string }
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Message != "Hello, Alice!" {
		t.Fatalf("expected greeting to include Alice, got %q", body.Message)
	}
}

func TestGetGreetingDefaultsWhenNameMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body struct{ Message string }
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Message != "Hello, World!" {
		t.Fatalf("expected default greeting, got %q", body.Message)
	}
}

func TestGetGreetingDefaultsWhenNameEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)
	rec := httptest.NewRecorder()
	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body struct{ Message string }
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Message != "Hello, World!" {
		t.Fatalf("expected default greeting, got %q", body.Message)
	}
}
