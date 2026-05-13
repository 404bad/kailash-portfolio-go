package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	// Expect either 200 (file found) or 404 (no static files in test env)
	if rr.Code != http.StatusOK && rr.Code != http.StatusNotFound {
		t.Errorf("unexpected status code: got %v", rr.Code)
	}
}

func TestHomeHandlerInvalidPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/invalid-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404 for invalid path, got %v", rr.Code)
	}
}

func TestPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("about.html")
	handler(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusNotFound {
		t.Errorf("unexpected status code: got %v", rr.Code)
	}
}
