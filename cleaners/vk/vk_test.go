package vk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOKResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"response": 1}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.Clear("http://etokavkaz.ru")
	if err != nil {
		t.Error("http://etokavkaz.ru must be reset ok")
	}

	if !result {
		t.Error("Result must be true")
	}
}

func TestErrorResponse(t *testing.T) {
	client := New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"response": 0}`)
	}))
	defer ts.Close()

	client.BaseURL = ts.URL

	result, err := client.Clear("http://etokavkaz.ru")
	if err == nil {
		t.Error("http://etokavkaz.ru must be not reset")
	}

	if result {
		t.Error("Result must be false")
	}
}
