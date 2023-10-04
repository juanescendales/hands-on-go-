package main

import "fmt"

type myInterface interface {
	methodWithoutParameters()
}

type myStruct struct {
	
}

func (m myStruct) methodWithoutParameters(){
	fmt.Println("Hello World")
}

func main(){

	// You cannot instanciate an interface
	//x := myInterface{}

	a := myStruct{}
	a.methodWithoutParameters()

	// Verify that a is an instance of myInterface
	var b myInterface = a
	b.methodWithoutParameters()

	

}

