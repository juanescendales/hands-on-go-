// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func lenghtName(name string, lenghtMap map[string]int, wm *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	lenghtMap[name] = len(name)
	mu.Unlock()
	wm.Done()
}

func main() {
	// given a list of names
	names := []string{"Andrea", "Pedro", "Laura", "Juan"}
	// initialize a map to store the length of each name
	var lenghtMap map[string]int = make(map[string]int)
	// initialize a wait group for the goroutines that will be launched
	var wm sync.WaitGroup
	wm.Add(len(names))
	// launch a goroutine for each name we want to process
	var mu sync.Mutex
	for _, name := range names {
		go lenghtName(name, lenghtMap, &wm, &mu)
	}
	// wait for all goroutines to finish
	wm.Wait()
	// print the map
	fmt.Printf("%#v\n", lenghtMap)
}
