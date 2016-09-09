package facebook

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
		fmt.Fprintln(w, `{"url": "https://test.com/", "type": "website", "title": "Test"}`)
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
		fmt.Fprintln(w, `{"error": {"message": "An unknown error has occurred.", "type": "OAuthException", "code": 1, "fbtrace_id": "Gsh/g5x4dzY"}}`)
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
