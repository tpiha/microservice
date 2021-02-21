package main

type Payload struct {
	Ts       int64   `json:"ts,omitempty"`
	MetricID int     `json:"metric_id,omitempty"`
	Value    float64 `json:"value,omitempty"`
}
