package main

import (
	"sync"

	"github.com/go-macaron/binding"
	"go.uber.org/ratelimit"
	macaron "gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

var m *macaron.Macaron

var db *gorm.DB

var mm *MetricsManager

var wm *WorkerManager

var rl ratelimit.Limiter

var mu sync.Mutex

func main() {
	db = initDb()

	mm = initMetricsManager()

	wm = initWorkerManager()

	m = macaron.New()

	m.Use(macaron.Renderer())

	m.Post("/process", binding.Json(Payload{}), processPayload)

	m.Run("0.0.0.0", 8080)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/process", hello)

	// srv := &http.Server{
	// 	Addr:              ":8080",
	// 	ReadTimeout:       5 * time.Second,
	// 	ReadHeaderTimeout: 5 * time.Second,
	// 	WriteTimeout:      10 * time.Second,
	// 	IdleTimeout:       5 * time.Second,
	// 	Handler:           mux,
	// }
	// log.Println(srv.ListenAndServe())
	// log.Fatal(http.ListenAndServe(":8080", mux))
}
