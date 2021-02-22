package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	concurrent    = 100
	semaphoreChan = make(chan struct{}, concurrent)
	client        *resty.Client
	wg            sync.WaitGroup
)

func doWork(item int) {
	semaphoreChan <- struct{}{} // block while full
	go func(item int) {
		defer func() {
			<-semaphoreChan // read to release a slot
			wg.Done()
		}()

		resp, err := client.R().
			SetBody(Payload{Ts: 1600346834534, MetricID: 1, Value: 43.2}).
			SetResult(&Payload{}).
			Post(fmt.Sprintf("http://localhost:8080/process?v=%d", item))

		if err != nil ||
			resp.StatusCode() != 200 {
			log.Println(resp)
			log.Println(err)
			doWork(item)
		}

	}(item)
}

func checkHealth() bool {
	resp, err := client.R().
		Get(fmt.Sprintf("http://localhost:8080/health"))

	if err != nil {
		return false
	}

	if !strings.Contains(string(resp.Body()), "OK") {
		return false
	}

	return true
}

func checkHealthCount() bool {
	resp, err := client.R().
		Get(fmt.Sprintf("http://localhost:8080/health-count"))

	if err != nil {
		return false
	}

	if !strings.Contains(string(resp.Body()), "10000") {
		return false
	}

	return true
}

func TestMicroservice(t *testing.T) {
	client = resty.New()

	if checkHealth() {
		fmt.Printf("Testing microservice...\n")

		for i := 1; i <= 10000; i++ {
			wg.Add(1)
			doWork(i)
		}

		wg.Wait()

		wait := true

		for wait {
			wait = !checkHealthCount()
			fmt.Printf("Waiting for records...\n")
			time.Sleep(time.Second)
		}
	}
}
