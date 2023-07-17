// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	var numero float32 = 3.1425
	var entero uint16= uint16(numero)

	fmt.Printf("numero %v entero %v\n", numero, entero)
}
