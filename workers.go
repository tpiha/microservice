package main

import (
	"log"
	"time"
)

const TICK = 1

type WorkerManager struct {
	Jobs    chan *Payload
	LastJob *Payload
}

func (wm *WorkerManager) addDatapoint(p *Payload) {
	wm.Jobs <- p
}

func (wm *WorkerManager) process() {
	go func() {
		for {
			if len(wm.Jobs) >= 1000 {
				jobs := []*Payload{}
				for i := 0; i < 1000; i++ {
					job := <-wm.Jobs
					jobs = append(jobs, job)
				}
				wm.saveBatch(jobs)
				log.Println(len(jobs))
			}

			time.Sleep(time.Millisecond * TICK * 100)
		}
	}()
}

func (wm *WorkerManager) saveBatch(jobs []*Payload) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println(err)
	}

	diff := uint(0)

	for _, job := range jobs {
		var err error

		if wm.LastJob != nil {
			diff = wm.processDiff(job.Value, wm.LastJob.Value)
		}

		dp := &Datapoint{Timestamp: uint64(job.Ts), Metric: mm.TS, Value: uint(job.Value), Diff: uint(diff)}
		if err = tx.Create(dp).Error; err != nil {
			tx.Rollback()
			log.Println(err)
		}

		wm.LastJob = job
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err)
	}
}

func (wm *WorkerManager) processDiff(currentValue float64, lastValue float64) uint {
	cvi := uint(currentValue * float64(100))
	lvi := uint(lastValue * float64(100))
	return cvi - lvi
}

func initWorkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.Jobs = make(chan *Payload, 1000000)
	wm.process()
	return wm
}
