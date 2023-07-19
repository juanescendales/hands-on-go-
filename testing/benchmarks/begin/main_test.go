// testing/benchmarks/begin/main_test.go
package main

import "testing"



// write a benchmark for sum
func BenchmarkSum(b *testing.B) {
	var input []int = []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		sum(input...)
	}

}
func BenchmarkSumAny(b *testing.B) {
	var input []any = []any{1,2,3,4}
	for i := 0; i < b.N; i++ {
		sumAny(input)
	}
}
