// interfaces/empty/begin/main.go
package main
import "fmt"
// define whatAmI which takes in an argument of any type and returns its underlying value's type
func whatAmI(element any) string{
	return fmt.Sprintf("Type: %T", element)
}
func main() {
	// invoke whatAmI with an int
	fmt.Println(whatAmI(1))

	// invoke whatAmI with a string
	fmt.Println(whatAmI("hello"))
}
