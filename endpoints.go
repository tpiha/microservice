package main

import (
	"strconv"

	macaron "gopkg.in/macaron.v1"
)

func processPayload(ctx *macaron.Context, p Payload) {
	wm.addDatapoint(&p)
	ctx.JSON(200, map[string]string{"status": "success"})
}

func healthCheck(ctx *macaron.Context) string {
	return "OK"
}

func healthCountCheck(ctx *macaron.Context) string {
	var dps []*Datapoint
	db.Find(&dps)
	return strconv.Itoa(len(dps))
}
