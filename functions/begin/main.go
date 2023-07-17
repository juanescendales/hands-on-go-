// functions/begin/main.go
package main

import (
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Hello World"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	var greeting string = fmt.Sprint("Hello %s!", name)
	return greeting
}

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = fmt.Sprint("Hello %s !, your age is %s", name, strconv.Itoa(age))
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
//

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("Toni"))

	// invoke divide function
	// fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	// fmt.Println(divide(5, 0))
}
