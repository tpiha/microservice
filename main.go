package main

import (
	"sync"

	"github.com/go-macaron/binding"
	macaron "gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

var m *macaron.Macaron

var db *gorm.DB

var mm *MetricsManager

var wm *WorkerManager

var mu sync.Mutex

var mudb sync.Mutex

func main() {
	db = initDb()

	mm = initMetricsManager()

	wm = initWorkerManager()

	m = macaron.New()

	m.Use(macaron.Renderer())

	m.Post("/process", binding.Json(Payload{}), processPayload)
	m.Get("/health", healthCheck)
	m.Get("/health-count", healthCountCheck)

	m.Run("0.0.0.0", 8080)
}
