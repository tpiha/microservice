package main

import (
	"log"
	"net/http"

	macaron "gopkg.in/macaron.v1"
)

func processPayload(ctx *macaron.Context, p Payload) {
	counter++
	log.Println(counter)

	// wm.addDatapoint(p)
	ctx.JSON(200, map[string]string{"status": "success"})
}

func hello(w http.ResponseWriter, r *http.Request) {
	counter++
	log.Println(counter)
}
