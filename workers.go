package main

import "log"

const TICK = 1

var (
	concurrent    = 20
	semaphoreChan = make(chan struct{}, concurrent)
)

type WorkerManager struct {
	Jobs []*Payload
}

func (wm *WorkerManager) addDatapoint(dp Payload) {
	// go func() {
	// 	wm.Jobs <- &dp
	// 	log.Println(len(wm.Jobs))
	// }()
	wm.Jobs = append(wm.Jobs, &dp)

	log.Println(len(wm.Jobs))
}

func (wm *WorkerManager) process() {
	semaphoreChan <- struct{}{}

	go func() {
		defer func() {
			<-semaphoreChan // read to release a slot
		}()

		// p := <-wm.Jobs
		// go wm.savePayload(p)
		// time.Sleep(time.Second * TICK)

		wm.process()
	}()
}

func (wm *WorkerManager) savePayload(p *Payload) {
	dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
	if err := db.Create(dp).Error; err != nil {
		log.Println(err)
		wm.savePayload(p)
	}
}

func initWowkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.process()
	return wm
}
