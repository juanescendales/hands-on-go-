// testing/table-driven/begin/main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, -1},
		{"three", []int{1, 2, 3}, 6},
	}

	// range over the tests and run them as subtests
	for _, test := range tests {
		testFunction := func(t *testing.T){
			if sum(test.input...) != test.want {
				t.Errorf("Error in sum test case %v + %v != %v", test.input[0], test.input[1], test.want)
			}
		}
		t.Run(test.name, testFunction)
	}
}


func TestSumFail(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		undesirable  int
	}{
		{"one", []int{1}, 0},
		{"two", []int{1, -2}, 3},
		{"three", []int{1, 2, 3}, 7},
	}

	// range over the tests and run them as subtests
	for _, test := range tests {
		testFunction := func(t *testing.T){
			if sum(test.input...) == test.undesirable {
				t.Errorf("Error in sum test case %v + %v != %v", test.input[0], test.input[1], test.undesirable)
			}
		}
		t.Run(test.name, testFunction)
	}
}
