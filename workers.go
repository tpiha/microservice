package main

import (
	"log"
	"time"
)

const TICK = 1

var (
	concurrent    = 1000
	semaphoreChan = make(chan struct{}, concurrent)
)

type WorkerManager struct {
	Jobs []*Payload
}

func (wm *WorkerManager) addDatapoint(dp *Payload) {
	mu.Lock()
	defer mu.Unlock()
	wm.Jobs = append(wm.Jobs, dp)
}

func (wm *WorkerManager) process() {
	go func(wm *WorkerManager) {
		for {
			if len(wm.Jobs) >= 1000 {
				mu.Lock()
				jobs := wm.Jobs[:1000]
				wm.Jobs = wm.Jobs[1000:]
				mu.Unlock()
				log.Println(len(jobs))
				log.Println(len(wm.Jobs))

				// dpl := []*Datapoint{}

				for _, p := range jobs {
					go func(p *Payload) {
						dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
						mudb.Lock()
						defer mudb.Unlock()
						if err := db.Create(dp).Error; err != nil {
							log.Println(err)
						}
					}(p)
				}
			} else {
				log.Println(len(wm.Jobs))
				time.Sleep(time.Second * TICK)
			}
		}
	}(wm)
}

func (wm *WorkerManager) savePayload(pd []*Payload) {
	dpl := []*Datapoint{}

	for _, p := range pd {
		dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
		dpl = append(dpl, dp)
	}
	if err := db.CreateInBatches(dpl, 10).Error; err != nil {
		log.Println(err)
		wm.savePayload(pd)
	}
}

func initWorkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.process()
	return wm
}
