package main

import "testing"

type testBase struct {
	name  string
	input string
}

type testCase struct {
	testBase
	want int
}

type TestCounter struct {
	counterName string
	testCases   []testCase
	counter     counter
}

func (tc *TestCounter) test(t *testing.T) {
	for _, test := range tc.testCases {
		testFunction := func(t *testing.T) {
			if tc.counter.count(test.input) != test.want {
				t.Errorf("Error with letter counter %v: count(%v) != %v\n", tc.counterName, test.input, test.want)
			}
		}
		t.Run(test.name, testFunction)
	}
}

func createTestCases(testBases []testBase, wantList ...int) []testCase {
	if len(testBases) != len(wantList) {
		panic("Not same lenght between want list and test bases")
	}

	testCases := []testCase{}
	for i := 0; i < len(testBases); i++ {
		testCases = append(testCases, testCase{
			testBase: testBases[i],
			want:     wantList[i],
		})
	}
	return testCases
}
