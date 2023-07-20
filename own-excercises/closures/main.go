package main

import "fmt"

func main() {
	a := 1
	fmt.Println("### FIRST EXPERIMENT ###")
	func() {
		var b int = a // Is a copy
		b += 1

		var c int = a
		fmt.Println("c:", c)
	}()

	fmt.Println("a:", a)

	fmt.Println("### SECOND EXPERIMENT ###")
	func() {
		var b int = a // Is a copy
		b += 1

		var c int = a
		c = 4
		fmt.Println("c:", c)
	}()

	fmt.Println("a:", a)

	fmt.Println("### THIRD EXPERIMENT ###")
	func() {
		a += 1 // Is a direct modification to the reference.

		var c int = a
		c = 4
		fmt.Println("c:", c)
	}()

	fmt.Println("a:", a)

}
