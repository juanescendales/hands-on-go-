package main

import "fmt"

func main() {
	var a int = 1
	b := &a

	var c int = *b
	c = 2

	fmt.Println(a, *b, c)
}
