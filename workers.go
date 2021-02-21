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
	wm.Jobs = append(wm.Jobs, dp)
}

func (wm *WorkerManager) process() {
	if len(wm.Jobs) >= 1000 {
		// var wg sync.WaitGroup
		// semaphoreChan <- struct{}{}
		// wg.Add(1)
		// go func(wg sync.WaitGroup) {
		// 	defer func() {
		// 		<-semaphoreChan // read to release a slot
		// 		wg.Done()
		// 		wm.process()
		// 	}()

		// }(wg)
		// wg.Wait()
		// wm.savePayload(wm.Jobs[:1000])
		// wm.Jobs = wm.Jobs[1001:]
		// wm.process()
		time.Sleep(time.Second * TICK)
		log.Println(len(wm.Jobs))
		wm.process()
	} else {
		time.Sleep(time.Second * TICK)
		wm.process()
	}
}

func (wm *WorkerManager) savePayload(pd []*Payload) {
	dpl := []*Datapoint{}

	for _, p := range pd {
		dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
		dpl = append(dpl, dp)
	}
	if err := db.Create(dpl).Error; err != nil {
		log.Println(err)
		wm.savePayload(pd)
	}
}

func initWorkerManager() WorkerManager {
	wm := WorkerManager{}
	go wm.process()
	return wm
}
