// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Example
	Example()
	// declare a signal channel to read from
	signalChan := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		signalChan <- struct{}{}
	}()
	// use a default case in select to prevent blocking
	for {
		select {
		case <-signalChan:
			fmt.Println("Signal channel")
			return
		default:
			fmt.Println("No signal")
			time.Sleep(200 * time.Millisecond)
		}
	}

}
