// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"
func main() {
	// create a local variable "val" and assign it an int value
	var val int
	// print the value of the local variable "val"
	fmt.Println(val)
	// print the value of the package-level variable "val"
	printGlobalVal()
	// update the package-level variable "val" and print it again
	updateGlobalVal("GloVal")
	printGlobalVal()
	// print the pointer address of the local variable "val"
	var pointer *int
	pointer = &val
	fmt.Println(pointer)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*pointer = 123
	fmt.Println(val)

}

func printGlobalVal(){
	fmt.Println(val)
}

func updateGlobalVal(newValue string){
	val = newValue
}