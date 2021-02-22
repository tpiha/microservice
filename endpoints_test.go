package main

import (
	"os"
	"testing"

	"gopkg.in/macaron.v1"
)

func TestHealtchCheck(t *testing.T) {
	ctx := &macaron.Context{}
	response := healthCheck(ctx)
	if response != "OK" {
		t.Errorf("got %s, want %s", response, "OK")
	}
}

func TestHealtchCountCheck(t *testing.T) {
	os.Remove("microservice.db")
	db = initDb()
	ctx := &macaron.Context{}
	response := healthCountCheck(ctx)
	if response != "0" {
		t.Errorf("got %s, want %s", response, "0")
	}
}
