package main

import "log"

type MetricsManager struct {
	TS *Metric
}

func initMetricsManager() *MetricsManager {
	mm := &MetricsManager{}

	mm.TS = &Metric{Name: "timestamp"}
	if err := db.FirstOrCreate(mm.TS, mm.TS).Error; err != nil {
		log.Printf("[initMetricsManager] %s", err)
	}

	return mm
}
