package main

import (
	"time"

	"github.com/go-macaron/binding"
	"github.com/martini-contrib/throttle"
	macaron "gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

var m *macaron.Macaron

var db *gorm.DB

var mm *MetricsManager

var wm *WorkerManager

var counter uint64

func main() {
	counter = 0

	db = initDb()

	mm = initMetricsManager()

	wm = initWowkerManager()

	m = macaron.New()

	m.Use(macaron.Renderer())

	m.Use(throttle.Policy(&throttle.Quota{
		Limit:  1,
		Within: time.Microsecond,
	}))

	m.Post("/process", binding.Json(Payload{}), processPayload)

	// tollbooth.LimitFuncHandler(tollbooth.NewLimiter(100000, nil)),

	m.Run("0.0.0.0", 8080)
}
