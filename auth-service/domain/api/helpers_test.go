package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadJSON(t *testing.T) {
	app := &AuthHandler{}
	jsonData := `{"key":"value"}`
	req, err := http.NewRequest("POST", "/", bytes.NewBufferString(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var data map[string]string
	err = app.readJSON(rr, req, &data)
	if err != nil {
		t.Fatal(err)
	}

	if val, ok := data["key"]; !ok || val != "value" {
		t.Fatalf("readJSON did not decode JSON correctly: got %+v, want %s", data, "value")
	}

	oversizeJSON := `{"key":"` + string(make([]byte, 1048577)) + `"}`
	req, err = http.NewRequest("POST", "/", bytes.NewBufferString(oversizeJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	err = app.readJSON(rr, req, &data)
	if err == nil {
		t.Fatal("Expected error for oversize JSON, got nil")
	}

	multiJSON := `{"key":"value"}{"key":"value"}`
	req, err = http.NewRequest("POST", "/", bytes.NewBufferString(multiJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	err = app.readJSON(rr, req, &data)
	if err == nil {
		t.Fatal("Expected error for multiple JSON values, got nil")
	}
}

func TestWriteJSON(t *testing.T) {
	app := &AuthHandler{}
	w := httptest.NewRecorder()

	data := map[string]string{
		"message": "test message",
	}

	err := app.writeJSON(w, http.StatusOK, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"message":"test message"}`
	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", w.Body.String(), expected)
	}

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", w.Code, http.StatusOK)
	}

	if value := w.Header().Get("Content-Type"); value != "application/json" {
		t.Errorf("unexpected header: got %v want %v", value, "application/json")
	}
}

func TestErrorJSON(t *testing.T) {
	app := &AuthHandler{}
	w := httptest.NewRecorder()

	err := app.errorJSON(w, errors.New("test error"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"error":true,"message":"test error"}`
	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", w.Body.String(), expected)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected status code: got %v want %v", w.Code, http.StatusBadRequest)
	}

	if value := w.Header().Get("Content-Type"); value != "application/json" {
		t.Errorf("unexpected header: got %v want %v", value, "application/json")
	}
}
