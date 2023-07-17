// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type

func whatAmI(element any) string {
	switch element.(type){
	case int:
		return fmt.Sprint("This element has an int type")
	case string:
		return fmt.Sprint("This element has a string type")
	default:
		return fmt.Sprintf("This element has an unsoported type %T", element)
	}
}
func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
