package main

const TICK = 1

type WorkerManager struct {
	Jobs chan *Payload
}

func (wm *WorkerManager) addDatapoint(dp Payload) {
	go func() { wm.Jobs <- &dp }()
}

func (wm *WorkerManager) process() {
	go func() {
		for i := 0; i < 1000; i++ {
			p := <-wm.Jobs
			go wm.savePayload(p)
		}
		wm.process()
	}()
}

func (wm *WorkerManager) savePayload(p *Payload) {
	dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
	db.Create(dp)
}

func initWowkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.Jobs = make(chan *Payload)
	wm.process()
	return wm
}
