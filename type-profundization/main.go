package main

import "fmt"

type specialInt int
type superSpecialInt int  // type definition
type cuteInt = specialInt // alias declaration

func (si specialInt) double() specialInt {
	return si * 2
}
func main() {
	var a specialInt = 1
	// var b superSpecialInt = 2
	var c cuteInt = 12

	fmt.Println(a.double())
	// fmt.Println( b.double()) compilation error
	fmt.Println(c.double())
}
