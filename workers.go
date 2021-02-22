package main

import (
	"log"
	"time"
)

const TICK = 1

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
			if len(wm.Jobs) >= 100 {
				mu.Lock()
				jobs := wm.Jobs[:100]
				wm.Jobs = wm.Jobs[100:]
				mu.Unlock()
				log.Println(len(jobs))
				log.Println(len(wm.Jobs))

				// sql := "INSERT into `datapoints` (timestamp, metric_id) VALUES "

				// for _, p := range jobs {
				// 	sql += fmt.Sprintf("(%d, %d),", p.Ts, mm.TS.ID)
				// }

				// sql = strings.TrimSuffix(sql, ",")

				// db.Exec(sql)

				// dpl := []*Datapoint{}
				// var err error

				// for _, p := range jobs {
				// 	dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
				// 	dpl = append(dpl, dp)
				// }

				// if err = db.CreateInBatches(dpl, 100).Error; err != nil {
				// 	log.Println(err)
				// }

				// check := true
				// for check {
				// 	dplnew := []*Datapoint{}
				// 	for _, dp := range dpl {
				// 		if dp.ID == 0 {
				// 			dplnew = append(dplnew, dp)
				// 		}
				// 	}

				// 	if len(dplnew) == 0 {
				// 		check = false
				// 	} else {
				// 		log.Println(len(dplnew))
				// 		if err = db.CreateInBatches(dplnew, 100).Error; err != nil {
				// 			log.Println(err)
				// 		}
				// 	}
				// }

				for _, p := range jobs {
					go func(p *Payload) {
						var err error
						dp := &Datapoint{Timestamp: uint64(p.Ts), Metric: mm.TS}
						mudb.Lock()
						defer mudb.Unlock()
						if err = db.Create(dp).Error; err != nil {
							log.Println(err)
						}
						for err != nil || dp.ID == 0 {
							if err = db.Create(dp).Error; err != nil {
								log.Println(err)
							}
						}
					}(p)
				}
			} else {
				time.Sleep(time.Second * TICK)
			}
		}
	}(wm)
}

func initWorkerManager() *WorkerManager {
	wm := &WorkerManager{}
	wm.process()
	return wm
}
