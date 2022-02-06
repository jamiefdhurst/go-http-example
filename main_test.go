package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(index)

	handler.ServeHTTP(rr, req)

	expected := "Hi there!"
	if rr.Body.String() != expected {
		t.Errorf("Method returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
