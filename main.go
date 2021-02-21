package main

import (
	"github.com/go-macaron/binding"
	macaron "gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

var m *macaron.Macaron

var db *gorm.DB

var mm *MetricsManager

var wm *WorkerManager

func main() {
	db = initDb()

	mm = initMetricsManager()

	wm = initWowkerManager()

	m = macaron.New()

	m.Use(macaron.Renderer())

	m.Post("/process", binding.Json(Payload{}), processPayload)

	m.Run("0.0.0.0", 8080)
}
