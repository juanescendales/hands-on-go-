// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	Speak()
	Walk()
}
// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) Speak() {
	fmt.Printf("Hi'm speaking ...\n")
}
func (p person) Walk(){
	fmt.Printf("I'm walking ...\n")
}
// implement the Stringer interface for the person type
func (p person) String() string{
	return fmt.Sprintf("Hi my name is %v\n", p.name)
}
// define a dog type that can walk but not speak
type dog struct {
	name string
}
func (d dog) Walk(){
	fmt.Printf("I'm walking guau guau...")
}
func main() {
	p := person{
		name : "Juan",
	}
	// invoke with a person
	doHumanThings(p)

	// can we invoke with a dog?
	// doHumanThings(dog{
	// 	name : "Zeus",
	// })

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.Speak()
	h.Walk()
}
