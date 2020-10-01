// Copyright 2020 Business Process Technologies. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ChFunc interface {
	Echo() error // echoes channel
}

type ChEchoer struct {
	Chan chan int
}

func NewChEchoer(c chan int) ChFunc {
	return &ChEchoer{Chan: c}
}

func (c *ChEchoer) Echo() error {
	fmt.Printf("echo %v \n", <-c.Chan)
	return nil
}

type chanMap map[chan int]ChFunc
type chanPayload chan int

const (
	payloads = 3
)

func main() {
	done := make(chan bool)
	stop := make(chan interface{})
	wrkCh := make(chanPayload, 5)
	resCh := make(chanPayload, 5)

	defer close(wrkCh)
	go func(wrk chanPayload, res chanPayload, done chan<- bool, stop <-chan interface{}) {

	loopFor:
		for {
			select {
			case i := <-wrk:
				fmt.Printf("processed: i=%v \n", i)
				res <- i
			case <-stop:
				fmt.Println("closing")
				break loopFor
			}
		}

		done <- true
		fmt.Println("done")
	}(wrkCh, resCh, done, stop)

	go func(res chanPayload, stop chan<- interface{}) {
		defer close(stop)
		processed := 0

		for {
			select {
			case i := <-res:
				fmt.Printf("worked: i=%v \n", i)
				processed++
				if processed == payloads {
					return
				}
			}
		}
	}(resCh, stop)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < payloads; i++ {
		wrkCh <- rand.Intn(20)
	}

	<-done
}
