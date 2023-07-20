package main

import (
	"fmt"
	"time"
)

func Example() {
	dataChannel := make(chan int)
	controlChannel := make(chan struct{})

	// Producer: Send data to dataChannel indefinitely
	go func() {
		i := 1
		for {
			dataChannel <- i
			i++
			time.Sleep(time.Second) // Simulate data production
		}
	}()

	// Consumer: Process data from dataChannel and check for the control signal
	go func() {
		for {
			select {
			case data := <-dataChannel:
				fmt.Println("Received:", data)
				// Process data here

			case <-controlChannel:
				fmt.Println("Control signal received. Breaking out of the loop.")
				return // Exit the goroutine, breaking the loop
			}
		}
	}()

	// Let the program run for some time, and then send the control signal to break the loop.
	time.Sleep(5 * time.Second)
	fmt.Println("Sending control signal to stop the loop.")
	controlChannel <- struct{}{} // Send the control signal

	// Wait for a while to see the "Control signal received" message
	time.Sleep(2 * time.Second)
}
