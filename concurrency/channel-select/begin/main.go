// concurrency/channel-select/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare an empty struct channel for signaling
	signalCh := make(chan struct{})
	// declare a timer channel
	timeCh := time.After(2 * time.Second)
	// launch a goroutine to signal after 1 second
	go func(signalCh chan<- struct{}) {
		time.Sleep(1 * time.Second)
		signalCh <- struct{}{}
	}(signalCh)
	// wait for a signal on either channel
	select {
	case <-signalCh:
		fmt.Println("Signal channel")
	case <-timeCh:
		fmt.Println("Time channel")
	}
}
