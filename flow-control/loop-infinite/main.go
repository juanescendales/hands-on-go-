// flow-control/loop-infinite/main.go
package main

import "fmt"

func main() {
	i := 0
	for {
		if i == 7{
			fmt.Println(i)
			break
		}
		i++
		fmt.Println("---")
	}
}
