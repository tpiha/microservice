package main

import (
	"strconv"

	macaron "gopkg.in/macaron.v1"
)

func processPayload(ctx *macaron.Context, p Payload) {
	wm.addDatapoint(&p)
	ctx.JSON(200, map[string]string{"status": "success"})
}

func healthCountCheck(ctx *macaron.Context) string {
	var dps []*Datapoint
	db.Find(&dps)
	return strconv.Itoa(len(dps))
}
