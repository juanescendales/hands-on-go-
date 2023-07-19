// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count

var testBases []testBase = []testBase{
	{name: "%!8dsSA89", input: "%!8dsSA89"},
	{name: "%ds SA su", input: "%ds SA su"},
	{name: "129381sda", input: "129381sda"},
}

func TestLetterCounter(t *testing.T) {
	var testCases []testCase = createTestCases(testBases, 4, 6, 3)
	counter := letterCounter{}
	testCounter := TestCounter{
		counterName: "letter",
		testCases:   testCases,
		counter:     counter,
	}
	testCounter.test(t)
}

// write a test for numberCounter.count
func TestNumberCounter(t *testing.T) {
	var testCases []testCase = createTestCases(testBases, 3, 0, 6)
	counter := numberCounter{}
	testCounter := TestCounter{
		counterName: "number",
		testCases:   testCases,
		counter:     counter,
	}
	testCounter.test(t)
}

// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T) {
	var testCases []testCase = createTestCases(testBases, 2, 3, 0)
	counter := symbolCounter{}
	testCounter := TestCounter{
		counterName: "symbol",
		testCases:   testCases,
		counter:     counter,
	}
	testCounter.test(t)
}
