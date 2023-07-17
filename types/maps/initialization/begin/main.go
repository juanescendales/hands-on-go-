// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["jc"] = author{
		name: "Juan Esteban Cendales",
	}
	authors["jac"] = author{
		name: "Jaime Andres Cendales",
	}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	autor := authors["jc"]
	fmt.Printf("%#v\n", autor.name)
}
