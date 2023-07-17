// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var a any = "1"
	fmt.Println(a.(string))
	// perform a type assertion and handle the error
	if _, valid := a.(int); !valid {
		fmt.Println("Error of assertion")
	}
}
