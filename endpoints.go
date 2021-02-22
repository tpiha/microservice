package main

import (
	"net/http"

	macaron "gopkg.in/macaron.v1"
)

func processPayload(ctx *macaron.Context, p Payload) {
	wm.addDatapoint(&p)
	ctx.JSON(200, map[string]string{"status": "success"})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// counter++
	// log.Println(counter)
	w.Write([]byte("OK"))
}
