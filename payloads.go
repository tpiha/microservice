package main

type Payload struct {
	Ts       int64   `json:"ts"`
	MetricID int     `json:"metric_id"`
	Value    float64 `json:"value"`
}
