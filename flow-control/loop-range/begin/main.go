// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	var slice []int = []int{10, 20, 30, 40, 50}

	// use for-range to iterate over the slice and print each value
	for i, n := range slice {
		fmt.Println(i, ":", n)

	}

	// declare a map of strings to ints
	var dict map[string]int = map[string]int{"a" : 1, "b" : 2, "c" : 3}

	// use for-range to iterate over the map and print each key/value pair
	for key, v := range dict {
		fmt.Println(key, ":", v)
	}
}
