// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum
func TestSum(t *testing.T) {
	want := 4
	input := []int{1, 3}
	if sum(input...) != want {
		t.Error("Error with sum function")
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("### SETUP ###")
	m.Run()
	fmt.Println("### TEARDOWN ###")
}
