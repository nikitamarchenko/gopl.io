/*

Exercise 9.5: Write a program with two goroutines that send messages back and
forth over two unbuffered channels in ping-pong fashion. How many
communications per second can the program sustain?

Result: average 2.5m ops/sec

*/

package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	durationFlag := flag.String("d", "10s", "duration in time.Duration format")
	flag.Parse()
	duration, err := time.ParseDuration(*durationFlag)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	done := make(chan struct{})
	var c uint
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		loop:
		for {
			select {
			case v := <-ch1:
				c++
				ch2 <- v
			case <-done:
				close(ch2)
				break loop
			}
		}
		for range ch1 {}
	}()
	go func() {
		defer wg.Done()
		for v := range ch2 {
			ch1 <- v
		}
		close(ch1)
	}()
	w := time.After(duration)
	ch1 <- struct{}{}
	<-w
	done <- struct{}{}
	wg.Wait()
	fmt.Printf("Op/sec %d\ntotal communications %d\ntotal duration %s\n", c/uint(duration.Seconds()), c, duration)
}
