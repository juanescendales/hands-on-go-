// concurrency/channel-ranging/begin/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fill a channel with ints up to max
func fill(ch chan<- int, max int) {
	// ensure non-deterministic random numbers
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	// loop and fill the channel up to its capacity qwith random numbersyy
	for i := 0; i < cap(ch); i++ {
		ch <- random.Intn(max)
	}
	// close the channel and signal that no more values will be sent
	close(ch)
}

func main() {
	// create a channel with a buffer size
	ch := make(chan int, 5)
	// invoke the fill function as a goroutine
	go fill(ch, 100)
	// range over the channel and print out the values
	for number := range ch {
		fmt.Println(number)
	}
}
