package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initSignalHandler() (chan os.Signal, chan struct{}, chan struct{}) {
	sigs := make(chan os.Signal, 1)
	mcrdone := make(chan struct{})
	dbdone := make(chan struct{})

	signal.Notify(sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	go func() {
		<-sigs
		log.Println("Shutting down gracefully...")

		// Shut down web server first
		mcrdone <- struct{}{}

		// Finish with background tasks
		for {
			if len(wm.Jobs) == 0 {
				dbdone <- struct{}{}
			}
			time.Sleep(time.Second * TICK)
		}
	}()

	return sigs, mcrdone, dbdone
}
