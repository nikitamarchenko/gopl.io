/*

Exercise 9.4: Construct a pipeline that connects an arbitrary number of
goroutines with channels. What is the maximum number of pipeline stages you can
create without running out of memory? How long does a value take to transit the
entire pipeline?

*/

package main

import (
	"fmt"
	"time"
)

// max = 12325415
func maxPipelines() (result uint) {
	enter := make(chan struct{})
	from := enter
	for {
		to := make(chan struct{})
		go func(f chan struct{}, t chan struct{}) {
			d := <-f
			t <- d
		}(from, to)
		result++
		from = to
		println(result)
	}
}

func pipelines(count uint) (enter chan struct{}, exit chan struct{}) {
	enter = make(chan struct{})
	from := enter
	for i := 0; i < int(count); i++ {
		exit = make(chan struct{})
		go func(f chan struct{}, t chan struct{}) {
			d := <-f
			t <- d
		}(from, exit)
		from = exit
	}
	return
}

// on 12215415 17.455959076s
func main() {
	//maxPipelines()
	f, t := pipelines(12325415 - 110000)
	fmt.Println("start")
	now := time.Now()
	f <- struct{}{}
	<-t
	fmt.Println(time.Since(now))
}
