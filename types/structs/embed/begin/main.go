// types/structs/embed/begin/main.go
package main

import "fmt"

type person struct {
	first string
	last  string
}

// fullName returns the full name of a person
func (p person) fullName() string {
	return p.first + " " + p.last
}

// define author and embed person
type author struct {
	person
	signature string
}

// override fullName method for author
func (a author) fullName() string {
	return a.first + " " + a.last + " signature: " + a.signature
}

func main() {
	// initialize and print a person's full name
	p := person{
		first: "Toni",
		last:  "Morrison",
	}
	fmt.Println(p.fullName())

	// initialize and print an author's full name
	a := author{
		person: person{
			first: "Juan",
			last:  "Cendales",
		},
		signature: "juanescendales",
	}
	fmt.Println(a.fullName())
}
