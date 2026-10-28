package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestItems(t *testing.T) {
	srv := httptest.NewServer(newStore().routes())
	defer srv.Close()

	res, _ := http.Post(srv.URL+"/items", "application/json", strings.NewReader(`{"name":"buku"}`))
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("create: got %d", res.StatusCode)
	}
	res, _ = http.Get(srv.URL + "/items/1")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("get: got %d", res.StatusCode)
	}
	res, _ = http.Post(srv.URL+"/items", "application/json", strings.NewReader(`{}`))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("empty name: got %d", res.StatusCode)
	}
	res, _ = http.Get(srv.URL + "/items/99")
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("missing: got %d", res.StatusCode)
	}
}
