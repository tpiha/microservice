package main

import (
	"log"
	"time"

	"gorm.io/gorm"
)

const TICK = 1

type WorkerManager struct {
	Jobs chan *Payload
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
				wm.saveBatch(jobs, *db, *mm)
			}

			time.Sleep(time.Millisecond * TICK * 100)
		}
	}()
}

func (wm *WorkerManager) saveBatch(jobs []*Payload, dbo gorm.DB, mmo MetricsManager) {
	tx := dbo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println(err)
	}

	for _, job := range jobs {
		var err error
		dp := &Datapoint{Timestamp: uint64(job.Ts), Metric: mmo.TS}
		if err = tx.Create(dp).Error; err != nil {
			tx.Rollback()
			log.Println(err)
		}
	}

	mudb.Lock()
	defer mudb.Unlock()
	if err := tx.Commit().Error; err != nil {
		log.Println(err)
	}
}

func initWorkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.Jobs = make(chan *Payload, 1000000)
	wm.process()
	return wm
}
