package main

import (
	"os"
	"testing"

	"gopkg.in/macaron.v1"
)

func TestHealtchCountCheck(t *testing.T) {
	os.Remove("microservice.db")
	db = initDb()
	ctx := &macaron.Context{}
	response := healthCountCheck(ctx)
	if response != "0" {
		t.Errorf("got %s, want %s", response, "0")
	}
}
