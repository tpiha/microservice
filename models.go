package main

import "gorm.io/gorm"

type Metric struct {
	gorm.Model
	Name string `sql:"size:255;unique_index"`
}

type Datapoint struct {
	gorm.Model
	MetricID  uint
	Metric    *Metric
	Timestamp uint64
	Value     uint
}
