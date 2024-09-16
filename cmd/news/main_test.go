package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHomePage(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Home page content"))
    })

    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("expected status OK, got %v", rr.Code)
    }

    expected := "Home page content"
    if rr.Body.String() != expected {
        t.Errorf("expected body %v, got %v", expected, rr.Body.String())
    }
}