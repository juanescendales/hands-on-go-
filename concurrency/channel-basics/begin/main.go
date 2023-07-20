// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, channel chan <- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	channel <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	channel := make(chan int)
	// invoke the sum function as a goroutine
	go sum(nums, channel)

	// force main thread to sleep
	result:= <- channel
	fmt.Println(result)
}
