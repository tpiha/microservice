package main

import (
	macaron "gopkg.in/macaron.v1"
)

func processPayload(ctx *macaron.Context, p Payload) {
	ctx.JSON(200, map[string]string{"status": "success"})
}
