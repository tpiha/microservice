package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	concurrentTest    = 100
	semaphoreChanTest = make(chan struct{}, concurrentTest)
	client            *resty.Client
	wg                sync.WaitGroup
)

func doWork(item int) {
	semaphoreChanTest <- struct{}{} // block while full
	go func(item int) {
		defer func() {
			<-semaphoreChanTest // read to release a slot
			wg.Done()
		}()

		resp, err := client.R().
			SetBody(Payload{Ts: 1600346834534, MetricID: 1, Value: 43.2}).
			SetResult(&Payload{}).
			Post(fmt.Sprintf("http://localhost:8080/process?v=%d", item))
		// if err != nil ||
		// 	!strings.Contains(string(resp.Body()), "success") ||
		// 	resp.StatusCode() != 200 {
		// 	log.Println(resp)
		// 	log.Println(err)
		// 	doWork(item)
		// }

		if err != nil ||
			resp.StatusCode() != 200 {
			log.Println(resp)
			log.Println(err)
			doWork(item)
		}

		log.Println("done")

	}(item)
}

func TestMicroservice(t *testing.T) {
	fmt.Printf("Testing microservice...\n")

	client = resty.New()

	client.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})

	for i := 1; i <= 50000; i++ {
		wg.Add(1)
		doWork(i)
	}

	wg.Wait()
}
