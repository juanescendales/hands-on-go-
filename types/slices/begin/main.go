// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "Jaime")
	names = append(names, "Mary")
	names = append(names, "Juanes")

	// print the slice
	fmt.Println(names)
	// slice the slice using syntax slice[low:high]
	fmt.Println(names[:2])
}
